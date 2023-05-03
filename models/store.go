package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	AreaID  uint      `json:"areaid" form:"areaid"`
	Address string    `json:"address" form:"address"`
	Phone   string    `json:"phone" form:"phone"`
	Email   string    `json:"email" form:"email"`
	Cleaner []Cleaner `gorm:"foreignKey:CleanerID"`
}
