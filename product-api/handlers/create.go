package handlers

import (
	"fmt"
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
)

func (p *products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(ProductKey{}).(*data.Product)

	p.l.Printf("[DEBUG] inserting product %#v\n", prod)
	data.AddProduct(prod)

	rw.Header().Add("Content-Type", "application/json; charset=utf-8")
	rw.Header().Add("Locator", fmt.Sprintf("%v/%v", r.RequestURI, prod.ID))
	rw.WriteHeader(http.StatusCreated)
}
