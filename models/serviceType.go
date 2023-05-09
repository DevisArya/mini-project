package models

type ServiceType struct {
	Id                 uint                `json:"Id" form:"Id" gorm:"primarykey"`
	Name               string              `json:"Name" form:"Name" validate:"required"`
	Price              uint64              `json:"Price" form:"Price" validate:"required"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:ServiceTypeID"`
}
type ServiceTypeResponse struct {
	Id    uint
	Name  string
	Price uint64
}
