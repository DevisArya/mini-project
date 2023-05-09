package models

type Admin struct {
	Id       uint   `json:"Id" form:"Id" gorm:"primarykey"`
	Name     string `json:"Name" form:"Name" validate:"required"`
	Address  string `json:"Address" form:"Address" validate:"required"`
	Phone    string `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Role     bool
	Password string `json:"password" form:"password" validate:"required"`
}
type AdminResponse struct {
	Id      uint
	Name    string
	Address string
	Phone   string
	Email   string
	Role    bool
}
type AdminResponseLogin struct {
	Id    int
	Name  string
	Email string
	Token string
}
