package main

import (
	"log"

	"github.com/iuryfukuda/auction/api"
)

func main() {
	server := api.NewServer()
	log.Fatal(server.Start(":3000"))
}
