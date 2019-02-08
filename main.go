package main

import (
	"os"

	_ "github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	bkr "github.com/micro/go-plugins/broker/grpc"
	cli "github.com/micro/go-plugins/client/grpc"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	srv "github.com/micro/go-plugins/server/grpc"
	"github.com/micro/micro/cmd"
	_ "github.com/micro/go-plugins/selector/static"

	// disable namespace by default
	"github.com/micro/micro/api"
)

func main() {
	plugin.Register(cors.NewPlugin())
	// disable namespace
	api.Namespace = ""

	// set values for registry/selector
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	os.Setenv("MICRO_SELECTOR", "static")

	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()
	

	// init command
	cmd.Init()
}
