package main

import (
	"fmt"
	"net/http"
)

/*
The logger is a middleware that wraps all handlers.
When a request is received, it will log some meaningful data.

This only works for incoming requests, and not outgoing responses.
*/
func logger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request has been received")
	}

	return http.HandlerFunc(fn)
}
