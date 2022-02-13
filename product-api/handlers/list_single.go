package handlers

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/proto"
	"github.com/marcosvieirajr/go-multi-tier-microservices/product-api/data"
)

func (p *products) HandleListSingle() http.HandlerFunc {
	type request struct{}
	type response struct {
		ID          int     `json:"id"`
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float32 `json:"price" validate:"gt=0"`
		SKU         string  `json:"sku" validate:"required,sku"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		id := getProductID(r)
		p.log.Debug("get record with", "id", id)

		prod, err := data.GetProductById(id)
		switch err {
		case nil:
		case data.ErrProductNotFound:
			p.log.Error("fetching product", "error", err)
			p.respond(rw, r, GenericError{Message: err.Error()}, http.StatusNotFound)
			return
		default:
			p.log.Error("fetching product", "error", err)
			p.respond(rw, r, GenericError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		// get exchange rate
		in := &proto.RateRequest{
			// Base:        proto.Currencies(proto.Currencies_value["EUR"]),
			// Destination: proto.Currencies(proto.Currencies_value["USD"]),
			Base:        "EUR",
			Destination: "GBP",
		}
		out, err := p.cc.GetRate(context.Background(), in)
		if err != nil {
			p.log.Error("error getting rate", "error", err)
			p.respond(rw, r, GenericError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		p.log.Info("Resp ", hclog.Fmt("%#v", out))
		prod.Price = prod.Price * out.Rate

		p.respond(rw, r, prod, http.StatusOK)
	}
}
