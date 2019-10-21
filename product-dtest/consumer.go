package main

import (
	"fmt"
	"secondsKill/product-dtest/common"
	"secondsKill/product-dtest/rabbitmq"
	"secondsKill/product-dtest/repositories"
	"secondsKill/product-dtest/services"
)

func main()  {
	db,err:=common.NewMysqlConn()
	if err !=nil {
		fmt.Println(err)
	}
	//创建product数据库操作实例
	product := repositories.NewProductManager("product",db)
	//创建product serivce
	productService:=services.NewProductService(product)
	//创建Order数据库实例
	order := repositories.NewOrderMangerRepository("order",db)
	//创建order Service
	orderService := services.NewOrderService(order)

	rabbitmqConsumeSimple :=rabbitmq.NewRabbitMQSimple("yyccQQuProduct")
	rabbitmqConsumeSimple.ConsumeSimple(orderService,productService)
}
