package main

import "secondsKill/RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQRouting("exImooc","imooc_two")
	imoocOne.RecieveRouting()
}
