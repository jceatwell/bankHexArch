package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jceatwell/bankHexArch/domain"
	"github.com/jceatwell/bankHexArch/service"
)

// type Customer struct {
// 	ID      int    `json:"id" xml:"id,attr"`
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	Zipcode string `json:"zip_code" xml:"zip""`
// }

// CustomerHandlers : Concrete Implementation(s) of REST Service
type CustomerHandlers struct {
	service service.CustomerService
}

// getAllCustomers : Handler function for /customers route
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{1, "John", "Port Elizabeth", "1234"},
	// 	{2, "Rob", "Durban", "12222"},
	// }

	customers, _ := ch.service.GetAllCustomer()
	log.Println(customers)

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		customerList := struct {
			XMLName   xml.Name           `xml:"customers"`
			Customers *[]domain.Customer `xml:"list>customer"`
		}{Customers: &customers}

		xml.NewEncoder(w).Encode(customerList)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
