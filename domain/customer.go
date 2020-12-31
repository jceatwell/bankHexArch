package domain

import "github.com/jceatwell/bankHexArch/errs"

// Customer : Define Business Object
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// CustomerRepository : Repository interface which sits on the edge of the Domain
// Note: Anthing defining this protocol should be able to connect to this port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
