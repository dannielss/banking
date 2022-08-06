package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/dannielss/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Daniel", City: "Guarulhos", ZipCode: "03130310"},
	// 	{Name: "Daniel2", City: "Guarulhos", ZipCode: "03130310"},
	// }

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")

		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(customers)
	}
}
