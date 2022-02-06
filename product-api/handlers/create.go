package handlers

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/marcosvieirajr/go-multi-tier-microservices/product-api/data"
)

func (p *products) HandleCreate() http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		// fetch the product from the context
		prod := r.Context().Value(ProductKey{}).(*data.Product)

		p.log.Debug("inserting product", "product", hclog.Fmt("%#v", prod))
		data.AddProduct(prod)

		rw.Header().Add("Locator", fmt.Sprintf("%v/%v", r.RequestURI, prod.ID))
		p.respond(rw, r, nil, http.StatusCreated)
	}
}
