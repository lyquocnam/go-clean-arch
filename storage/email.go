package storage

import (
"clean/model"
"context"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/postgres"
)

type emailStorage struct {
	db *gorm.DB
}

func NewEmailStorage(db *gorm.DB) *emailStorage {
	return &emailStorage{db: db}
}

type EmailStorage interface {
	GetEmailsByCustomerId(ctx context.Context, customerId int) ([]*model.Email, error)
}

func (s *emailStorage) GetEmailsByCustomerId(ctx context.Context, customerId int) ([]*model.Email, error) {
	var result []*model.Email
	err := s.db.Find(&result, "customer_id = ?", customerId).Error
	return result, err
}
