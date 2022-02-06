package handlers

import (
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/product-api/data"
)

func (p *products) HandleListAll() http.HandlerFunc {
	type request struct{}
	type response []struct {
		ID          int     `json:"id"`
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float32 `json:"price" validate:"gt=0"`
		SKU         string  `json:"sku" validate:"required,sku"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		p.log.Debug("get all records")

		// fetch the products from the datastore
		prods := data.GetProducts()

		p.respond(rw, r, prods, http.StatusOK)
	}
}
