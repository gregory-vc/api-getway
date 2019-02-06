package main

import (
	_ "github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
)

func main() {
	plugin.Register(cors.NewPlugin())
	cmd.Init()
}
