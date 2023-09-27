package model

type GetCurrentUsersProfileResponse struct {
	Country     string `json:"country"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	ID          string `json:"id"`
	URI         string `json:"uri"`
}

type GetFollowedArtistsResponse struct {
	Artists struct {
		ArtistList []Artist `json:"items"`
		Next       string   `json:"next"`
		Total      int      `json:"total"`
		Cursors    struct {
			After string `json:"after"`
		} `json:"cursors"`
		Limit int    `json:"limit"`
		HREF  string `json:"href"`
	} `json:"artists"`
}
