package dto

type UserDetailResponse struct {
	ID           int    `json:"id"`
	NisnOrNip    string `json:"nisn_or_nip"`
	Fullname     string `json:"fullname"`
	Nickname     string `json:"nickname"`
	DateOfBirth  string `json:"date_of_birth"`
	PlaceOfBirth string `json:"place_of_birth"`
	Email        string `json:"email"`
	Address      string `json:"address"`
}
