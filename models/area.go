package models

type Area struct {
	Id          uint          `json:"Id" form:"Id" gorm:"primarykey"`
	Name        string        `json:"Name" form:"Name" validate:"required"`
	Store       []Store       `gorm:"foreignKey:AreaID"`
	Transaction []Transaction `gorm:"foreignKey:AreaID"`
}
type AreaResponse struct {
	Id    uint
	Name  string
	Store []Store
}
type AreaResponseStore struct {
	Id    uint
	Name  string
	Store []StoreResponse
}
type AreaResponseCreate struct {
	Id   uint
	Name string
}
