package main

import (
	"log"
	"net/http"
	"github.com/iuryfukuda/auction/app"
	"github.com/iuryfukuda/auction/handles"
)

func main() {
	mux := http.NewServeMux()

	a := app.New()
	mux.HandleFunc("/bid", handles.Bid)
	// mux.HandleFunc("/stats", handles.Stats)
	log.Printf("a=(%+v)", a)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
