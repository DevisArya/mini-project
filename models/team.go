package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name        string        `json:"name" form:"name"`
	Transaction []Transaction `gorm:"foreignKey:TeamID"`
	Cleaner     []Cleaner     `gorm:"foreignKey:TeamID"`
}
