package dto

type UserResponse struct {
	ID         int    `json:"id"`
	Identifier string `json:"identifier"`
	RoleId     int    `json:"role_id"`
}
