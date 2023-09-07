package model

type GetCurrentUsersProfileResponse struct {
	Country     string `json:"country"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	ID          string `json:"id"`
	URI         string `json:"uri"`
}
