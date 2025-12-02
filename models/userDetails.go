package models

// import "time"

type UserDetails struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	UserId       int    `json:"user_id" gorm:"unique"`
	Fullname     string `json:"fullname"`
	Nickname     string `json:"nickname"`
	DateOfBirth  string `json:"date_of_birth"`
	PlaceOfBirth string `json:"place_of_birth"`
	Email        string `json:"email"`
	Address      string `json:"address"`
}

func (UserDetails) TableName() string {
	return "user_details"
}
