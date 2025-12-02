package dto

type UserDetailResponse struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id"`
	NisnOrNip    string `json:"nisn_or_nip"`
	Fullname     string `json:"fullname"`
	Nickname     string `json:"nickname"`
	DateOfBirth  string `json:"date_of_birth"`
	PlaceOfBirth string `json:"place_of_birth"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
}
