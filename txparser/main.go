package main

import (
	"log"
	"net/http"
)

// Port http server listen on.
const portNum string = ":8181"

func main() {

	// Spinning up the server.
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
