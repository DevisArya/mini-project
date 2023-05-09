package models

type Store struct {
	Id      uint      `json:"Id" form:"Id" gorm:"primarykey"`
	AreaId  uint      `json:"AreaId" form:"AreaId" validate:"required"`
	Address string    `json:"Address" form:"Address" validate:"required"`
	Phone   string    `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	Email   string    `json:"Email" form:"Email" validate:"required,email"`
	Cleaner []Cleaner `gorm:"foreignKey:StoreId"`
}
type StoreResponse struct {
	Id      uint
	AreaId  uint
	Address string
	Phone   string
	Email   string
}
