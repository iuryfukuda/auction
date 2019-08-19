package main

import (
	"log"

	"github.com/iuryfukuda/auction/api"
)

func main() {
	server := api.NewServer()
	log.Fatal(server.Run(":3000"))
}
