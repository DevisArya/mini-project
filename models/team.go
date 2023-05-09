package models

type Team struct {
	Id          uint          `json:"Id" form:"Id" gorm:"primarykey"`
	Name        string        `json:"Name" form:"Name" validate:"required"`
	Transaction []Transaction `gorm:"foreignKey:TeamID"`
	Cleaner     []Cleaner     `gorm:"foreignKey:TeamID"`
}

type TeamResponse struct {
	Id      uint
	Name    string
	Cleaner []Cleaner
}
type TeamResponseCreate struct {
	Id   uint
	Name string
}
