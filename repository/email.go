package repository

import (
	"clean/model"
	"clean/storage"
	"context"
)

type emailRepository struct {
	emailStorage storage.EmailStorage
}

func NewEmailRepository(emailStorage storage.EmailStorage) *emailRepository {
	return &emailRepository{emailStorage: emailStorage}
}

type EmailRepository interface {
	GetEmailsByCustomerId(ctxt context.Context, customerId int) ([]*model.Email, error)
}

func (repo *emailRepository) GetEmailsByCustomerId(ctx context.Context, customerId int) ([]*model.Email, error) {
	return repo.emailStorage.GetEmailsByCustomerId(ctx, customerId)
}