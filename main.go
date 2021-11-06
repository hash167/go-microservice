package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/main/handlers"
	"time"

	"github.com/gobuffalo/envy"
)

var bindAddress = envy.Get("BIND_ADDRESS", ":9090")

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)
	ph := handlers.NewProducts(l)

	//Create a new http Server Multiplexer
	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)
	sm.Handle("/products/", ph)

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
