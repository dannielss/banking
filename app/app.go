package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dannielss/banking/domain"
	"github.com/dannielss/banking/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	envVariables := []string{"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASS", "DB_ADDR", "DB_PORT", "DB_NAME"}

	for _, v := range envVariables {
		if os.Getenv(v) == "" {
			log.Fatal("Environment variable not defined... ", v)
		}
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
