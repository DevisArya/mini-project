package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	Phone       string `json:"phone" form:"phone" validate:"required,min=10,max=13"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Role        bool
	Password    string        `json:"password" form:"password" validate:"required"`
	Transaction []Transaction `gorm:"foreignKey:CustomerID"`
}

type CustomerResponse struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
