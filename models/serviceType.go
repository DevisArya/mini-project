package models

import "gorm.io/gorm"

type ServiceType struct {
	gorm.Model
	Name               string              `json:"name" form:"name" validate:"required"`
	Price              uint64              `json:"price" form:"price" validate:"required"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:ServiceTypeID"`
}
type ServiceTypeResponse struct {
	gorm.Model
	Name  string `json:"name" form:"name" validate:"required"`
	Price uint64 `json:"price" form:"price" validate:"required"`
}
