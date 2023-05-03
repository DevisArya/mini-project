package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	TransactionID uint `json:"transactionid" form:"transactionid"`
	ServiceTypeID uint `json:"servicetypeid" form:"servicetypeid"`
	Qty           uint `json:"qty" form:"qty"`
	TotalPrice    uint `json:"totalprice" form:"totalprice"`
}
