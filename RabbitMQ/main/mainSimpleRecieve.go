package main

import "secondsKill/RabbitMQ"

func main() {
	robbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	robbitmq.ConsumeSimple()
}
