package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jceatwell/bankHexArch/domain"
	"github.com/jceatwell/bankHexArch/logger"
	"github.com/jceatwell/bankHexArch/service"
)

func sanityChecks() {
	envVars := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, v := range envVars {
		if os.Getenv(v) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", v))
		}
	}
}

// Start starts the Server listening in at port 8080
func Start() {
	sanityChecks()

	// Define a Multiplexer
	router := mux.NewRouter()

	// Wiring (Define Service Instance and Repository Instance)

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// Start server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
