package main

import (
	"github.com/mrtechit/backend-assignment/internal/service"
	"log"
	"net/http"
)

// Port http server listen on.
const portNum string = ":8181"

func main() {

	http.HandleFunc("/block", service.GetCurrentBlock)
	http.HandleFunc("/subscribe", service.Subscribe)
	http.HandleFunc("/transactions", service.GetTransactions)

	// Spinning up the server.
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
