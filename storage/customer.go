package storage

import (
	"clean/model"
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type customerStorage struct {
	db *gorm.DB
}

func NewCustomerStorage(db *gorm.DB) *customerStorage {
	return &customerStorage{db: db}
}

type CustomerStorage interface {
	GetCustomerById(ctx context.Context, customerId int) (*model.Customer, error)
}

func (s *customerStorage) GetCustomerById(ctx context.Context, customerId int) (*model.Customer, error) {
	var result model.Customer
	err := s.db.First(&result, "id = ?", customerId).Error
	return &result, err
}