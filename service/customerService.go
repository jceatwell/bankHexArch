package service

import (
	"github.com/jceatwell/bankHexArch/domain"
	"github.com/jceatwell/bankHexArch/dto"
	"github.com/jceatwell/bankHexArch/errs"
)

// CustomerService : REST Port interface
type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService : This is the Business Logic Domain
// Impletments Service Interface and has a dependency on the Repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer : Receiver function for getAllCutomers
func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, len(customers))
	for i, c := range customers {
		response[i] = c.ToDto()
	}
	return response, nil

}

// GetCustomer : Receiver function for getCustomer
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

// NewCustomerService : factory helper funvyin to create default Customer REST Service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
