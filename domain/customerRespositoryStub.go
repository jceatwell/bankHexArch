package domain

// CustomerRepositoryStub : Stub adaptor for Server Side
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll : Implementation of CustomerRepository interface
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub : Responsible for creating all dummy customers
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "John", "Pretoria", "110011", "2000-01-01", "1"},
		{"1002", "Dave", "Cape Toen", "110012", "2000-01-02", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
