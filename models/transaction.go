package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CustomerID         uint `json:"customerid" form:"customerid" validate:"required"`
	TeamID             uint `json:"teamid" form:"teamid" validate:"required"`
	PaymentID          uint `json:"paymentid" form:"paymentid" validate:"required"`
	Status             bool
	Location           string `json:"location" form:"location" validate:"required"`
	AreaID             uint   `json:"areaid" form:"areaid" validate:"required"`
	Rating             uint
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID"`
}

type TransactionUpdateRating struct {
	Rating uint `json:"rating" form:"rating" validate:"required"`
}
type TransactionUpdateStatus struct {
	Status bool `json:"status" form:"status" validate:"required"`
}
