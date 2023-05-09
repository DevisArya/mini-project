package models

type Transaction struct {
	Id                 uint `json:"Id" form:"Id" gorm:"primarykey"`
	CustomerId         uint `json:"CustomerId" form:"CustomerId" validate:"required"`
	TeamId             uint `json:"TeamId" form:"TeamId" validate:"required"`
	PaymentId          uint `json:"PaymentId" form:"PaymentId" validate:"required"`
	Status             bool
	Location           string `json:"Location" form:"Location" validate:"required"`
	AreaId             uint   `json:"AreaId" form:"AreaId" validate:"required"`
	Rating             uint
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
type TransactionResponse struct {
	Id         uint
	CustomerId uint
	TeamId     uint
	PaymentId  uint
	Location   string
	AreaId     uint
}
type TransactionUpdateRating struct {
	Rating uint `json:"Rating" validate:"required"`
}
type TransactionUpdateStatus struct {
	Status bool `json:"Status" form:"status" validate:"required"`
}
