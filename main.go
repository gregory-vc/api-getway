package main

import (
	"os"

	_ "github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
	"github.com/micro/go-micro"

	// _ "github.com/micro/go-plugins/registry/kubernetes"
	stan "github.com/micro/go-plugins/broker/stan"
	"github.com/micro/go-micro/broker"
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
	cmd.Init(
		micro.Broker(
			stan.NewBroker(
				stan.ClusterID(os.Getenv("MICRO_BROKER_CLUSTER_ID")),
				broker.Addrs(os.Getenv("MICRO_BROKER_ADDRESS")),
			),
		),
	)
}
