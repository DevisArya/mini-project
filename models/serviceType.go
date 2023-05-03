package models

import "gorm.io/gorm"

type ServiceType struct {
	gorm.Model
	Name               string              `json:"name" form:"name"`
	Price              uint64              `json:"price" form:"price"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:ServiceTypeID"`
}
