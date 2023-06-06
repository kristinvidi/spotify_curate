package model

import (
	apptype "spotify_app/api/pkg/app_type"
)

type User struct {
	DisplayName string
	Email       string
	ID          apptype.UserID
	Country     apptype.UserCountryCode
}

type GetCurrentUsersProfileResponse struct {
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ExplicitContent struct {
		FilterEnabled bool `json:"filter_enabled"`
		FilterLocked  bool `json:"filter_locked"`
	} `json:"explicit_content"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		HREF  string `json:"href"`
		Total int    `json:"total"`
	}
	HREF    string  `json:"href"`
	ID      string  `json:"id"`
	Images  []Image `json:"images"`
	Product string  `json:"product"`
	Type    string  `json:"type"`
	URI     string  `json:"uri"`
}

func (g *GetCurrentUsersProfileResponse) User() *User {
	return &User{
		DisplayName: g.DisplayName,
		Email:       g.Email,
		ID:          apptype.UserID(g.ID),
		Country:     apptype.UserCountryCode(g.Country),
	}
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
