package main

import "secondsKill/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.*.two")
	imoocOne.RecieveTopic()
}
