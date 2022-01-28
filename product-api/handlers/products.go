package handlers

import (
	"log"
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
)

// Products is a http.Handler
type products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *products {
	return &products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler
// interface
func (p *products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getProducts(rw, r)
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func getProducts(rw http.ResponseWriter, r *http.Request) {
// getProducts returns the products from the data store

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func postProducts(rw http.ResponseWriter, r *http.Request) {
	log.Print("POST calledo")
	// ps := []product{}
	// json.Unmarshal([]byte(r.Body.Read()), &ps)
	// log.Println("****", ps)
}
