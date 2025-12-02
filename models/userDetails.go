package models

// import "time"

type UserDetails struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	UserId       int    `json:"user_id" gorm:"unique"`
	Fullname     string `json:"fullname" validate:"required"`
	Nickname     string `json:"nickname" validate:"required"`
	DateOfBirth  string `json:"date_of_birth" validate:"required"`
	PlaceOfBirth string `json:"place_of_birth" validate:"required"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address" validate:"required"`
}

func (UserDetails) TableName() string {
	return "user_details"
}
