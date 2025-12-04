package models

type Departments struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" validate:"max=100"`
	Nickname string `json:"nickname" validate:"max=50"`
	Slug     string `json:"slug" validate:"max=100"`
}
