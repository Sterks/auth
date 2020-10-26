package models

type Role struct {
	ID int `gorm:"primaryKey" json:"id"`
	Role string `json:"role"`
	Description string `json:"description"`
}
