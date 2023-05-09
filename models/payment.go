package models

type Payment struct {
	Id            uint          `json:"Id" form:"Id" gorm:"primarykey"`
	Name          string        `json:"Name" form:"Name" validate:"required"`
	PaymentType   string        `json:"PaymentType" form:"PaymentType" validate:"required"`
	PaymentNumber string        `json:"PaymentNumber" form:"PaymentNumber" validate:"required"`
	Transaction   []Transaction `gorm:"foreignKey:PaymentID"`
}

type PaymentResponse struct {
	Id            uint
	Name          string
	PaymentType   string
	PaymentNumber string
}
