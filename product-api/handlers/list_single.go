package handlers

import (
	"net/http"

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

		p.respond(rw, r, prod, http.StatusOK)
	}
}
