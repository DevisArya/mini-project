package models

type Payment struct {
	ID            uint          `json:"id" form:"id" gorm:"primarykey"`
	Name          string        `json:"name" form:"name" validate:"required"`
	PaymentType   string        `json:"paymentType" form:"paymentType" validate:"required"`
	PaymentNumber string        `json:"paymentNumber" form:"paymentNumber" validate:"required"`
	Transaction   []Transaction `gorm:"foreignKey:PaymentID"`
}

type PaymentResponse struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name" validate:"required"`
	PaymentType   string `json:"paymentType" form:"paymentType" validate:"required"`
	PaymentNumber string `json:"paymentNumber" form:"paymentNumber" validate:"required"`
}
