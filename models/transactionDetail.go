package models

type TransactionDetail struct {
	Id            uint `json:"Id" form:"Id" gorm:"primarykey"`
	TransactionId uint `json:"TransactionId" form:"TransactionId" validate:"required"`
	ServiceTypeId uint `json:"ServiceTypeId" form:"ServiceTypeId" validate:"required"`
	Qty           uint `json:"Qty" form:"Qty" validate:"required"`
	TotalPrice    uint `json:"TotalPrice" form:"TotalPrice" validate:"required"`
}
