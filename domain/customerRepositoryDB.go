package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

// FindAll : CustomerRepositoryDb, CustomerRepository.FindAll() implementation
func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSQL)
	if err != nil {
		log.Println("Error while querying customer tabl " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning customer " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// NewCustomerRepositoryDb : Factory Method
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// TODO: Need to bring in environmental variables
	client, err := sql.Open("mysql", "root:codepass@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
