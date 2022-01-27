package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(log *log.Logger) *hello {
	return &hello{l: log}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("handle Hello requests")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("error reading body", err)

		http.Error(rw, "unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Hello, %s", b)
}
