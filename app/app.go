package app

import (
	"log"
	"net/http"

	"github.com/dannielss/banking/domain"
	"github.com/dannielss/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
