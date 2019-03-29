package usecase

import (
	"clean/model"
	"clean/repository"
	"context"
	"time"
)

type customerUseCase struct {
	timeout time.Duration
	customerRepo repository.CustomerRepository
	emailRepo repository.EmailRepository
}

func NewCustomerUseCase(
	customerRepo repository.CustomerRepository,
	emailRepo repository.EmailRepository,
	timeout time.Duration) *customerUseCase {
	return &customerUseCase{
		customerRepo: customerRepo,
		emailRepo: emailRepo,
		timeout: timeout,
	}
}

func (u *customerUseCase) GetCustomerById(ctx context.Context, customerId int) (*model.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	customer, err := u.customerRepo.GetCustomerById(ctx, customerId)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, nil
	}

	emails, err := u.emailRepo.GetEmailsByCustomerId(ctx, customerId)
	if err != nil {
		return nil, err
	}
	if len(emails) < 1 {
		return nil, nil
	}
	customer.Emails = emails

	// logic...

	return customer, nil
}

type CustomerUseCase interface {
	GetCustomerById(ctx context.Context, customerId int) (*model.Customer, error)
}

