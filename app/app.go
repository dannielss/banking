package main

import (
	"log"
	"net/http"
)

func Start() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
