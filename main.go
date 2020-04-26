package main

import (
	"os"
	"log"
	"flag"
	"io/ioutil"
	"os/signal"

	"github.com/zbioe/auction/api"
	"github.com/zbioe/auction/db"
)

var dumpFilePath string
var port string

func main() {
	flag.Parse()

	server := api.NewServer()

	b, err := ioutil.ReadFile(dumpFilePath)
	if err == nil {
		bkp, err := db.MemFromJSON(b)
		if err == nil {
			server.Db = bkp
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		<-c
		b, err := server.Db.ToJSON()
		if err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(dumpFilePath, b, 0644); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}()

	log.Fatal(server.Run(port))
}

func init() {
	flag.StringVar(&dumpFilePath, "filepath", "auction.dump", "path for load and dump data")
	flag.StringVar(&port, "port", ":3000", "port for use the api")
}
