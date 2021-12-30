package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/main/handlers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gobuffalo/envy"
	"github.com/gorilla/mux"
)

var bindAddress = envy.Get("BIND_ADDRESS", ":9090")

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//Using gorilla mux

	sm := mux.NewRouter()

	ph := handlers.NewProducts(l)

	// register handler with the mux for GET requests
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct)

	// register middleware for docs

	opts := middleware.RedocOpts{SpecURL: "./swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//Post requests
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	//Put requests

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareValidateProduct)

	// Delete Requests

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	// http.ListenAndServe(":9090", sm)
	s := &http.Server{
		Addr:         bindAddress, //":9090"
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	// non blocking
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// Create a channel for signals
	sigChan := make(chan os.Signal)
	//  send the following signals to the channel when seen
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, gracefull shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
