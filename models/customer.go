package models

type Customer struct {
	Id          uint   `json:"Id" form:"Id" gorm:"primarykey"`
	Name        string `json:"Name" form:"Aame" validate:"required"`
	Address     string `json:"Address" form:"Address" validate:"required"`
	Phone       string `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	Email       string `json:"Email" form:"Email" validate:"required,email"`
	Role        bool
	Password    string        `json:"Password" form:"Password" validate:"required"`
	Transaction []Transaction `gorm:"foreignKey:CustomerID"`
}
type CustomerResponse struct {
	Id          uint
	Name        string
	Address     string
	Phone       string
	Email       string
	Transaction []Transaction `gorm:"foreignKey:CustomerID"`
}

type CustomerResponseLogin struct {
	Id    int
	Name  string
	Email string
	Token string
}
