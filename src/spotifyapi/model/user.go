package model

import "time"

type GetCurrentUsersProfileResponse struct {
	Country     string `json:"country"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	ID          ID     `json:"id"`
	URI         URI    `json:"uri"`
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

type TrackItem struct {
	AddedAt string          `json:"added_at"`
	Track   SimplifiedTrack `json:"track"`
}

func (t *TrackItem) SavedAt() time.Time {
	date, _ := time.Parse(time.RFC3339, t.AddedAt)
	return date
}

type GetUsersSavedTracksResponse struct {
	Items []TrackItem `json:"items"`
	Total int         `json:"total"`
}
