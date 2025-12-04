package models

type Users struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	Nip         string      `json:"nip" gorm:"unique" validate:"max=18"`
	Nisn        string      `json:"nisn" gorm:"unique" validate:"max=10"`
	Password    string      `json:"password" validate:"required"`
	RoleId      int         `json:"role_id" validate:"required"`
	Role        Roles       `json:"roles" gorm:"foreignKey:RoleId"`
	UserDetails UserDetails `json:"user_details" gorm:"foreignKey:UserID"`
}
