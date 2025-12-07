package models

type UserDetails struct {
	ID           int         `json:"id" gorm:"primaryKey"`
	UserId       int         `json:"user_id" gorm:"unique"`
	DepartmentId int         `json:"department_id" validate:"required"`
	Fullname     string      `json:"fullname" validate:"required"`
	Nickname     string      `json:"nickname" validate:"required"`
	DateOfBirth  string      `json:"date_of_birth" validate:"required"`
	PlaceOfBirth string      `json:"place_of_birth" validate:"required"`
	Email        string      `json:"email"`
	Phone        string      `json:"phone"`
	Address      string      `json:"address" validate:"required"`
	Department   Departments `json:"-" gorm:"foreignKey=DepartmentId"`
}

func (UserDetails) TableName() string {
	return "user_details"
}
