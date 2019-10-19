package repositories

import (
	"database/sql"
	"fmt"
	"secondsKill/product-dtest/common"
	"github.com/kataras/iris/core/errors"
	"secondsKill/product-dtest/datamodels"
	"strconv"
)

type IUserRepository interface {
	Conn() error
	Select(userName string) (user *datamodels.User, err error)
	Insert(user *datamodels.User) (userId int64, err error)
}
type UserManagerRepository struct {
	table string
	mysqlConn *sql.DB
}

func NewUserRepository(table string, db *sql.DB)  IUserRepository{
	return &UserManagerRepository{table, db}
}

func (u *UserManagerRepository) Conn() (err error) {
	if u.mysqlConn == nil {
		mysql, errMysql := common.NewMysqlConn()
		if errMysql != nil {
			return errMysql
		}
		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "user"
	}
	return
}

func (u *UserManagerRepository) Select(userName string) (user *datamodels.User, err error) {
	if userName == "" {
		return &datamodels.User{}, errors.New("条件不能为空！")
	}
	if err = u.Conn(); err != nil{
		return &datamodels.User{}, err
	}

	sql := "Select * from " + u.table + " where userName=?"
	fmt.Println(sql,"sql----")
	rows, errRows := u.mysqlConn.Query(sql, userName)
	defer rows.Close()
	if errRows != nil{
		return &datamodels.User{}, errRows
	}

	result := common.GetResultRow2(rows)
	fmt.Println(rows,"rows")
	if len(result) == 0{
		return &datamodels.User{}, errors.New("用户不存在！")
	}
	user = &datamodels.User{}
	fmt.Println(result,"result----")
	common.DataToStructByTagSql(result,user)
	return
}

func (u *UserManagerRepository) Insert(user *datamodels.User) (userId int64, err error) {
	if err = u.Conn(); err != nil {
		return
	}

	//sql := "INSERT " + u.table + " SET nickName=?,userName=?,passWord=?"
	//stmt, errStmt := u.mysqlConn.Prepare(sql)
	//if errStmt != nil {
	//	return userId, errStmt
	//}
	//result, errResult := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	//if errResult != nil {
	//	return userId, errResult
	//}

	sql :="INSERT INTO user (nickName,userName,passWord) values('%s','%s','%s')"
	result,err :=u.mysqlConn.Exec(fmt.Sprintf(sql,user.NickName, user.UserName, user.HashPassword))
	if err != nil{
		return 0,err
	}
	return result.LastInsertId()
}
func (u *UserManagerRepository) SelectByID(userId int64) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}
	sql := "select * from " + u.table + " where ID=" + strconv.FormatInt(userId, 10)
	row, errRow := u.mysqlConn.Query(sql)
	if errRow != nil {
		return &datamodels.User{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("用户不存在！")
	}
	user = &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return
}


