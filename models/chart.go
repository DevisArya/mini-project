package models

type Chart struct {
	Id            uint `json:"Id" form:"id" gorm:"primarykey"`
	CustomerId    uint `json:"CustomerId" form:"CustomerId" validate:"required"`
	ServiceTypeId uint `json:"ServiceTypeId" form:"ServiceTypeId" validate:"required"`
	Qty           uint `json:"Qty" form:"Qty" validate:"required"`
	TotalPrice    uint
}
type ChartResponse struct {
	Id            uint
	CustomerId    uint
	ServiceTypeId uint
	Qty           uint
	TotalPrice    uint
}
