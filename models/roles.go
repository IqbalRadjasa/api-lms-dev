package models

type Roles struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" validate:"max=100"`
}
