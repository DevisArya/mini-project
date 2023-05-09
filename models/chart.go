package models

import "gorm.io/gorm"

type Chart struct {
	gorm.Model
	CustomerID    uint `json:"customerId" form:"customerId" validate:"required"`
	ServiceTypeID uint `json:"serviceTypeId" form:"serviceTypeId" validate:"required"`
	Qty           uint `json:"qty" form:"qty" validate:"required"`
	TotalPrice    uint
}
type ChartResponse struct {
	Id            uint
	CustomerID    uint
	ServiceTypeID uint
	Qty           uint
	TotalPrice    uint
}
