package main

import (
	"log"
	"net/http"
	"github.com/iuryfukuda/auction/app"
	"github.com/iuryfukuda/auction/handle"
)

func main() {
	mux := http.NewServeMux()

	a := app.New()
	mux.HandleFunc("/bid", handle.Bid)
	// mux.HandleFunc("/stats", handle.Stats)
	log.Printf("a=(%+v)", a)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
