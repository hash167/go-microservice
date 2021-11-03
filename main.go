package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, error := ioutil.ReadAll(r.Body)
		if error != nil {
			http.Error(rw, "unable to read request body", http.StatusBadRequest)
			return
		}
		log.Printf("Data: %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("GoodBye World")
	})

	http.ListenAndServe(":9090", nil)

}
