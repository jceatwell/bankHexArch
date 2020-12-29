package service

import "github.com/jceatwell/bankHexArch/domain"

// CustomerService : REST Port interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// DefaultCustomerService : This is the Business Logic Domain
// Impletments Service Interface and has a dependency on the Repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer : Receiver function for getAllCutomers
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService : factory helper funvyin to create default Customer REST Service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
