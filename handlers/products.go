package handlers

import (
	"log"
	"net/http"
	"product-api/main/data"
	"regexp"
	"strconv"
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

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshall json", http.StatusBadRequest)
	}
	data.AddProduct(prod)

}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshall json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}

}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p.getProducts(rw)
		return
	}

	if r.Method == "POST" {
		p.addProduct(rw, r)
	}

	if r.Method == "PUT" {
		// Get ID from path
		reg := regexp.MustCompile("/([0-9])+")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URL, more than one id", g)
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URL, more than one capture group", g[0])
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}
		p.updateProduct(id, rw, r)

	}
	// Catch all
	// if no method is satisfied return error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
