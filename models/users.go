package models

type Users struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	Identifier  string      `json:"identifier" validate:"max=20"`
	Password    string      `json:"password" validate:"required"`
	RoleId      int         `json:"role_id" validate:"required"`
	Role        Roles       `json:"roles" gorm:"foreignKey:RoleId"`
	UserDetails UserDetails `json:"user_details" gorm:"foreignKey:UserID" validate:"-"`
}
