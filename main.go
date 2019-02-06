package main

import (
	_ "github.com/micro/go-plugins/broker/nats"
	"github.com/micro/micro/cmd"
)

func main() {
	cmd.Init()
}
