package handlers

import (
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
)

func (p *products) HandleUpdate() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := getProductID(r)
		p.log.Debug("updating record with", "id", id)

		// fetch the product from the context
		prod := r.Context().Value(ProductKey{}).(*data.Product)
		prod.ID = id

		err := data.UpdateProduct(*prod)
		if err == data.ErrProductNotFound {
			p.log.Error("product not found", "error", err)
			p.respond(rw, r, GenericError{Message: "Product not found in database"}, http.StatusNotFound)
			return
		}

		p.respond(rw, r, nil, http.StatusNoContent)
	}
}
