package usecase

import (
	"clean/mocks"
	"clean/model"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCustomerUseCase_GetCustomerById(t *testing.T) {
	mockCustomerRepo := &mocks.CustomerRepository{}
	mockEmailRepo := &mocks.EmailRepository{}
	customerUseCase := NewCustomerUseCase(mockCustomerRepo, mockEmailRepo, time.Duration(10 * time.Second))

	customerId := 1
	email1 := &model.Email{
		Address: "abc@gmail.com",
		CustomerID: uint(customerId),
	}
	emails := []*model.Email{
		email1,
	}
	expectCustomer := model.Customer{
		Model: gorm.Model{
			ID: uint(customerId),
		},
	}

	cases := []struct {
		name string
		want *model.Customer
		//actual *model.Customer
		err error
	}{
		{
			name: "Can not get customer",
			want: nil,
			err: errors.New("can not get customer"),
		},
		{
			name: "get customer successfully",
			want: &expectCustomer,
			err: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			mockCustomerRepo.On("GetCustomerById", mock.Anything, customerId).Return(c.want, c.err)
			mockEmailRepo.On("GetEmailsByCustomerId", mock.Anything, customerId).Return(emails, nil)
			actual, err := customerUseCase.GetCustomerById(context.TODO(), customerId)

			assert.Equal(t, c.want, actual)
			assert.Equal(t, c.err, err)

			mockCustomerRepo.AssertExpectations(t)
		})
	}
}
