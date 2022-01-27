package handlers

import (
	"log"
	"net/http"
)

type goodbye struct {
	l *log.Logger
}

func NewGoodBy(l *log.Logger) *goodbye {
	return &goodbye{l}
}

func (h *goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("handle Goodbye request")
	rw.Write([]byte("Goodbye"))
}
