package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	AreaID  uint      `json:"areaid" form:"areaid" validate:"required"`
	Address string    `json:"address" form:"address" validate:"required"`
	Phone   string    `json:"phone" form:"phone" validate:"required,min=10,max=13"`
	Email   string    `json:"email" form:"email" validate:"required,email"`
	Cleaner []Cleaner `gorm:"foreignKey:CleanerID"`
}
