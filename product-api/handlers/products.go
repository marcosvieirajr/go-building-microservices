package handlers

import (
	"log"
	"net/http"

	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
)

type products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *products {
	return &products{l}
}

func (p *products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getProducts(rw, r)
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
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
