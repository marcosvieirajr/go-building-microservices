package handlers

import (
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/product-api/data"
)

func (p *products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.log.Debug("deleting record by id", "id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.log.Error("deleting record id does not exist", "id", id)

		p.respond(rw, r, GenericError{Message: err.Error()}, http.StatusNotFound)
		return
	}

	if err != nil {
		p.log.Error("deleting record", "error", err)

		p.respond(rw, r, GenericError{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	p.respond(rw, r, nil, http.StatusNoContent)
}
