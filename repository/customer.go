package repository

import (
	"clean/model"
	"clean/storage"
	"context"
)

type customerRepository struct {
	customerStorage storage.CustomerStorage
}

func NewCustomerRepository(customerStorage storage.CustomerStorage) *customerRepository {
	return &customerRepository{customerStorage: customerStorage}
}

type CustomerRepository interface {
	GetCustomerById(ctxt context.Context, id int) (*model.Customer, error)
}

func (repo *customerRepository) GetCustomerById(ctxt context.Context, customerId int) (*model.Customer, error) {
	return repo.customerStorage.GetCustomerById(ctxt, customerId)
}