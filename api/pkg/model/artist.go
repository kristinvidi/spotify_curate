package model

import (
	apptype "spotify_app/api/pkg/app_type"
)

type Artist struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Genres apptype.Genres `json:"genres"`
}
