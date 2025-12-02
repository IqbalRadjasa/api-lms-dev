package dto

type UserResponse struct {
	ID     int    `json:"id"`
	Nisn   string `json:"nisn"`
	RoleId int    `json:"role_id"`
}
