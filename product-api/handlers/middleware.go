package handlers

import (
	"context"
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
)

func (p products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := p.decode(rw, r, prod)
		if err != nil {
			p.log.Error("deserializing product", "error", err)
			p.respond(rw, r, GenericError{Message: err.Error()}, http.StatusBadRequest)
			return
		}

		// validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.log.Error("validating product", "error", errs)
			// return the validation messages as an array
			p.respond(rw, r, ValidationError{Messages: errs.Errors()}, http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), ProductKey{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
