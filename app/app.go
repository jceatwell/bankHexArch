package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Start starts the Server listening in at port 8080
func Start() {

	// Define a Multiplexer
	router := mux.NewRouter()

	// This uses the default multiplexer
	router.HandleFunc("/greet", greetHandler).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	keys := make([]string, len(vars))
	for key := range vars {
		keys = append(keys, key)
	}
	log.Printf(strings.Join(keys, ","))
	fmt.Fprintf(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {

}
