package models

type Users struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Nisn     string `json:"nisn" gorm:"unique"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
}
