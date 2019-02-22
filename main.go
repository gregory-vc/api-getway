package main

import (
	_ "github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"

	// _ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/broker/stan"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/selector/static"
	_ "github.com/micro/go-plugins/transport/nats"
	"github.com/micro/micro/cmd"
)

func main() {
	plugin.Register(cors.NewPlugin())
	// disable namespace
	// api.Namespace = ""

	// set values for registry/selector
	// os.Setenv("MICRO_REGISTRY", "kubernetes")
	// os.Setenv("MICRO_SELECTOR", "static")

	// init command
	cmd.Init()
}
