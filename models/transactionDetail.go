package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	TransactionID uint `json:"transactionid" form:"transactionid" validate:"required"`
	ServiceTypeID uint `json:"servicetypeid" form:"servicetypeid" validate:"required"`
	Qty           uint `json:"qty" form:"qty" validate:"required"`
	TotalPrice    uint
}
