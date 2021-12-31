package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, error := ioutil.ReadAll(r.Body)
	if error != nil {
		http.Error(rw, "unable to read request body", http.StatusBadRequest)
		return
	}
	log.Printf("Data: %s\n", d)
	fmt.Fprintf(rw, "Hello %s", d)
}
