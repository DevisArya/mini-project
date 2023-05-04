package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	Phone       string `json:"phone" form:"phone"`
	Email       string `json:"email" form:"email"`
	Role        bool
	Password    string        `json:"password" form:"password"`
	Transaction []Transaction `gorm:"foreignKey:CustomerID"`
}

type CustomerResponse struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
