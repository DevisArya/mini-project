package models

type Admin struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required,min=10,max=13"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Role     bool
	Password string `json:"password" form:"password" validate:"required"`
}

type AdminResponse struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
