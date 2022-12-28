package main

import (
	"fmt"
	"log"
	"net/http"
)

const HOST = "localhost"
const PORT = 8081

type Config struct{}

func main() {
	ADDRESS := fmt.Sprintf("%s:%d", HOST, PORT)

	app := Config{}

	log.Printf("Starting the Broker service on port %d\n", PORT)

	// define the HTTP server
	server := &http.Server{
		Addr:    ADDRESS,
		Handler: app.routes(),
	}

	// start the server
	err := server.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}
