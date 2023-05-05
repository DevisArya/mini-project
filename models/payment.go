package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Name          string        `json:"name" form:"name" validate:"required"`
	PaymentType   string        `json:"paymenttype" form:"paymenttype" validate:"required"`
	PaymentNumber string        `json:"paymentnumber" form:"paymentnumber" validate:"required"`
	Transaction   []Transaction `gorm:"foreignKey:PaymentID"`
}
