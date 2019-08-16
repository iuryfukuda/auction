package main

import (
	"log"
	"net/http"
	"github.com/iuryfukuda/auction/app"
	"github.com/iuryfukuda/auction/handlers"
)

func main() {
	mux := http.NewServeMux()

	a := app.New()
	mux.HandleFunc("/bid", handlers.Bid)
	// mux.HandleFunc("/stats", handlers.Stats)
	log.Printf("a=(%+v)", a)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
