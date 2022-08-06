package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Daniel", City: "Guarulhos", ZipCode: "03130310"},
		{Name: "Daniel2", City: "Guarulhos", ZipCode: "03130310"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")

		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(customers)
	}
}
