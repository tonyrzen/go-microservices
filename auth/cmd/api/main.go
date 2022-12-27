package main

import (
	"auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const HOST = "localhost"
const PORT = 8080

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	ADDRESS := fmt.Sprintf("%s:%d", HOST, PORT)
	log.Printf("Starting the Broker service on port %d\n", PORT)

	// TODO connect to DB

	// set up config
	app := Config{}

	server := &http.Server{
		Addr:    ADDRESS,
		Handler: app.routes(),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
