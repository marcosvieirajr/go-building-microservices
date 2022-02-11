package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func New(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("handle RateRequest", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &currency.RateResponse{Rate: 0.5}, nil
}
