package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/proto"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/server"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.New(log)

	proto.RegisterCurrencyServer(gs, cs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
