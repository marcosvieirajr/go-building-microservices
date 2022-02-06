package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/marcosvieirajr/go-multi-tier-microservices/product-api/data"
)

// Products is a http.Handler
type products struct {
	log hclog.Logger
	v   *data.Validation
}

// NewProducts creates a products handler with the given logger
func NewProducts(l hclog.Logger, v *data.Validation) *products {
	return &products{log: l, v: v}
}

func (p *products) respond(rw http.ResponseWriter, r *http.Request, d interface{}, status int) {
	rw.Header().Add("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	if d != nil {
		err := json.NewEncoder(rw).Encode(d)
		if err != nil {
			p.log.Error("serializing product", "error", err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (p *products) decode(rw http.ResponseWriter, r *http.Request, d interface{}) error {
	return json.NewDecoder(r.Body).Decode(d)
}

type ProductKey struct{}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
