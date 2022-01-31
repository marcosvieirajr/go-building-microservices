package handlers

import (
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
)

func (p *products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	// fetch the products from the datastore
	prods := data.GetProducts()

	rw.Header().Add("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)

	// serialize the list to JSON
	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] get all records")
	}
}
