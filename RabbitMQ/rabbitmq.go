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
	rabbitmq := &RabbitMQ{
		QueueName:queueName,
		Exchange: exchange,
		Key: key,
		Mqurl: MQURL,
	}
	var err error
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
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


// 简单模式step： 1，创建简单模式下rabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq:= NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

// 2，简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	// 保证队里存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,//是否持久化
		false,//是否为自动删除
		false,//是否具有排他性
		false,//是否阻塞
		nil,////额外属性
	)
	if err!=nil {
		fmt.Println(err)
	}

	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据exchange 类型和 routerkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentEncoding: "text/plain",
			Body:[]byte(message),
		})

}

func (r *RabbitMQ) ConsumeSimple(){
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 接收消息
	msgs,err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,

	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 实现我们要处理的逻辑函数
			log.Printf("Received a message: %s", d.Body)
			fmt.Println(d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages, To exit press CTRL+C")
	<-forever

}


















