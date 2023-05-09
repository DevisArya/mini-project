package models

type Team struct {
	ID          uint          `json:"id" form:"id" gorm:"primarykey"`
	Name        string        `json:"name" form:"name" validate:"required"`
	Transaction []Transaction `gorm:"foreignKey:TeamID"`
	Cleaner     []Cleaner     `gorm:"foreignKey:TeamID"`
}

type TeamResponse struct {
	ID      uint      `json:"id" form:"id" gorm:"primarykey"`
	Name    string    `json:"name" form:"name" validate:"required"`
	Cleaner []Cleaner `gorm:"foreignKey:TeamID"`
}
