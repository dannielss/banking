package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/dannielss/banking/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")

		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
	} else {
		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Set("Content-Type", "application/xml")

			xml.NewEncoder(w).Encode(customer)
		} else {
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(customer)
		}
	}
}
