package main

import (
	"github.com/kataras/iris"
)

// 当前系统 和 服务器系统 不相同时、 需要交叉编译 => 当前文件目录下执行 目标系统：liunx， => GOOS=liunx GOARCH=amd64 go build productMain.go
// cdn测试 http://localhost:8082/html/htmlProduct.html

func main()  {
	//1.创建iris 实例
	app:=iris.New()

	//2.设置模板
	app.StaticWeb("/public","./web/public")
	//3.访问生成好的html静态文件
	app.StaticWeb("/html","./web/htmlProductShow")

	app.Run(
		iris.Addr("0.0.0.0:80"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
