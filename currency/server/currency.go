package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/proto"
)

type Currency struct {
	log hclog.Logger
	proto.UnimplementedCurrencyServer
}

func New(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *proto.RateRequest) (*proto.RateResponse, error) {
	c.log.Info("handle RateRequest", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &proto.RateResponse{Rate: 0.5}, nil
}
