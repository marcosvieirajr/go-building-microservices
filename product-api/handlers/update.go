package handlers

import (
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
)

func (p *products) Update(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("[DEBUG] updating record id", id)

	// fetch the product from the context
	prod := r.Context().Value(ProductKey{}).(*data.Product)
	prod.ID = id

	err := data.UpdateProduct(*prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
