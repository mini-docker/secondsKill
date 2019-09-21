package RabbitMQ

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
)

// url格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost
const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	Exchange  string // 交换机
	Key       string //
	Mqurl     string // 连接信息
}

// 创建RabbitMQ结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName:queueName,
		Exchange: exchange,
		Key: key,
		Mqurl: MQURL
	}
}

// 断开连接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err !=nil{
		log.Fatalf("%s:%s", message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq:=NewRabbitMQ(queueName,exchange:"",key:"")
	var err error
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}



