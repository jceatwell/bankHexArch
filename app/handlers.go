package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	ID      int    `json:"id" xml:"id,attr"`
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip""`
}

func greetHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{1, "John", "Port Elizabeth", "1234"},
		{2, "Rob", "Durban", "12222"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		customerList := struct {
			XMLName   xml.Name    `xml:"customers"`
			Customers *[]Customer `xml:"list>customer"`
		}{Customers: &customers}

		xml.NewEncoder(w).Encode(customerList)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
