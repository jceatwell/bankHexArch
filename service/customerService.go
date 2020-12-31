package service

import (
	"github.com/jceatwell/bankHexArch/domain"
	"github.com/jceatwell/bankHexArch/errs"
)

// CustomerService : REST Port interface
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService : This is the Business Logic Domain
// Impletments Service Interface and has a dependency on the Repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer : Receiver function for getAllCutomers
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}
	return s.repo.FindAll(status)
}

// GetCustomer : Receiver function for getCustomer
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// NewCustomerService : factory helper funvyin to create default Customer REST Service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
