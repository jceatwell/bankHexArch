package domain

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jceatwell/bankHexArch/errs"
	"github.com/jceatwell/bankHexArch/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// FindAll : CustomerRepositoryDb, CustomerRepository.FindAll() implementation
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error

	customers := make([]Customer, 0)

	if status == "" {
		findAllSQL := `select customer_id, name, city, zipcode, date_of_birth, status 
					   from customers`
		err = d.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL := `select customer_id, name, city, zipcode, date_of_birth, status 
					   from customers 
					   where status = ?`
		err = d.client.Select(&customers, findAllSQL, status)
	}

	if err != nil {
		logger.Error(fmt.Sprintf("Error while querying customer table %s", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSQL := `select customer_id, name, city, zipcode, date_of_birth, status 
					from customers 
					where customer_id = ?`

	var c Customer
	err := d.client.Get(&c, customerSQL, id)

	if err != nil {
		logger.Error(fmt.Sprintf("Error while scanning customer %s", err.Error()))
		switch err {
		case sql.ErrNoRows:
			return nil, errs.NewNotFoundError("Customer not found")
		default:
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

// NewCustomerRepositoryDb : Factory Method
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// TODO: Need to bring in environmental variables
	client, err := sqlx.Open("mysql", "root:codepass@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
