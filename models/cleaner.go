package models

type Cleaner struct {
	Id      uint   `json:"Id" form:"Id" gorm:"primarykey"`
	Name    string `json:"Name" form:"Name" validate:"required"`
	Address string `json:"Address" form:"Address" validate:"required"`
	Phone   string `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	Email   string `json:"Email" form:"Email" validate:"required,email"`
	TeamId  uint   `json:"TeamId" form:"TeamId" validate:"required"`
	StoreId uint   `json:"StoreId" form:"StoreId" validate:"required"`
}
