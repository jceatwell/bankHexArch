package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jceatwell/bankHexArch/domain"
	"github.com/jceatwell/bankHexArch/service"
)

// Start starts the Server listening in at port 8080
func Start() {

	// Define a Multiplexer
	router := mux.NewRouter()

	// Wiring (Define Service Instance and Repository Instance)
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
