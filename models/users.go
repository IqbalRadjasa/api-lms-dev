package models

type Users struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Nisn     string `json:"nisn" gorm:"unique" validate:"required,max=10"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"role_id" validate:"required"`
}
