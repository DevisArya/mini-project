package models

import "gorm.io/gorm"

type Cleaner struct {
	gorm.Model
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Phone   string `json:"phone" form:"phone"`
	Email   string `json:"email" form:"email"`
	TeamID  uint   `json:"teamid" form:"teamid"`
	StoreID uint   `json:"storeid" form:"storeid"`
}
