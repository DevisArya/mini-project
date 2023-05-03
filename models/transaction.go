package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CustomerID         uint                `json:"customerid" form:"customerid"`
	TeamID             uint                `json:"teamid" form:"teamid"`
	PaymentID          uint                `json:"paymentid" form:"paymentid"`
	TotalPrice         uint                `json:"totalprice" form:"totalprice"`
	Status             bool                `json:"status" form:"status"`
	Location           string              `json:"location" form:"location"`
	AreaID             uint                `json:"areaid" form:"areaid"`
	Rating             uint                `json:"rating" form:"rating"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
