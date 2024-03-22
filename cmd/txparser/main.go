package main

import (
	"github.com/mrtechit/backend-assignment/internal/service"
	"log"
	"net/http"
)

// Port http server listen on.
const portNum string = ":8181"

func main() {

	http.HandleFunc("v1/api/block", service.GetCurrentBlock)
	http.HandleFunc("v1/api/subscribe", service.Subscribe)
	http.HandleFunc("v1/api/transactions", service.GetTransactions)

	// Spinning up the server.
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
