package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/proto"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	cs := server.New(log)

	// register the currency server
	proto.RegisterCurrencyServer(gs, cs)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("unable to listen", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)
}
