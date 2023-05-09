package models

type Store struct {
	ID      uint      `json:"id" form:"id" gorm:"primarykey"`
	AreaID  uint      `json:"areaId" form:"areaId" validate:"required"`
	Address string    `json:"address" form:"address" validate:"required"`
	Phone   string    `json:"phone" form:"phone" validate:"required,min=10,max=13"`
	Email   string    `json:"email" form:"email" validate:"required,email"`
	Cleaner []Cleaner `gorm:"foreignKey:CleanerID"`
}
