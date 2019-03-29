package model

import "github.com/jinzhu/gorm"

type Email struct {
	gorm.Model
	Address string `json:"address"`
	CustomerID uint `json:"-"`
}
