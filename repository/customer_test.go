package repository

import (
	"clean/mocks"
	"clean/model"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerRepository_GetCustomerById(t *testing.T) {
	ctx := context.Background()
	customerStorage := &mocks.CustomerStorage{}
	customerRepo := NewCustomerRepository(customerStorage)

	customerId := 1
	expectCustomer := model.Customer{
		Model: gorm.Model{
			ID: uint(customerId),
		},
	}
	customerStorage.On("GetCustomerById", ctx, customerId).Return(&expectCustomer, nil)

	actual, err := customerRepo.GetCustomerById(ctx, customerId)
	assert.Equal(t, expectCustomer, *actual)
	assert.Nil(t, err)

}
