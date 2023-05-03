package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Name          string        `json:"name" form:"name"`
	PaymentType   string        `json:"paymenttype" form:"paymenttype"`
	PaymentNumber string        `json:"paymentnumber" form:"paymentnumber"`
	Transaction   []Transaction `gorm:"foreignKey:PaymentID"`
}
