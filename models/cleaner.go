package models

import "gorm.io/gorm"

type Cleaner struct {
	gorm.Model
	Name    string `json:"name" form:"name" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required,min=10,max=13"`
	Email   string `json:"email" form:"email" validate:"required,email"`
	TeamID  uint   `json:"teamid" form:"teamid" validate:"required"`
	StoreID uint   `json:"storeid" form:"storeid" validate:"required"`
}
