package models

import "gorm.io/gorm"

type Area struct {
	gorm.Model
	Name        string        `json:"name" form:"name"`
	Store       []Store       `gorm:"foreignKey:AreaID"`
	Transaction []Transaction `gorm:"foreignKey:AreaID"`
}
