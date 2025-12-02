package dto

type UserResponse struct {
	ID          int    `json:"id"`
	Nisn_or_Nip string `json:"nisn_or_nip"`
	RoleId      int    `json:"role_id"`
}
