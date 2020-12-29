package service

import "github.com/jceatwell/bankHexArch/domain"

// CustomerService : REST Port interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// DefaultCustomerService : Business domain repository dependency
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}
