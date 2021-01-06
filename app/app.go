package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jceatwell/bankHexArch/domain"
	"github.com/jceatwell/bankHexArch/logger"
	"github.com/jceatwell/bankHexArch/service"
	"github.com/jmoiron/sqlx"
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
	dbClient := getDbClient()
	customerRespositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRespositoryDb(dbClient)

	ch := CustomerHandlers{
		service: service.NewCustomerService(customerRespositoryDb),
	}
	ah := AccountHandler{
		service: service.NewAccountService(accountRepositoryDb),
	}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	// Start server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
