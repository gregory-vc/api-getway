package main

import (
	"log"

	"github.com/micro/go-micro"

	_ "github.com/micro/go-plugins/broker/nats"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
