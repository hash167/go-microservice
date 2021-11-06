package handlers

import (
	"log"
	"net/http"
	"product-api/main/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) getProducts(rw http.ResponseWriter) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()

	// Serialize list to JSON
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to martial json", http.StatusInternalServerError)
	}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p.getProducts(rw)
		return
	}
	// Catch all
	// if no method is satisfied return error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
