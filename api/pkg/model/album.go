package model

import (
	"time"

	apptype "spotify_app/api/pkg/app_type"
)

type ReleaseDatePrecision string

const (
	ReleaseDateYear  ReleaseDatePrecision = "year"
	ReleaseDateMonth ReleaseDatePrecision = "month"
	ReleaseDateDay   ReleaseDatePrecision = "day"
)

type Album struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	AlbumType string         `json:"album_type"`
	Genres    apptype.Genres `json:"genres"`
	Artists   Artists        `json:"artists"`
	TrackList struct {
		Tracks Tracks `json:"items"`
	} `json:"tracks"`
	TotalTracks          int                  `json:"total_tracks"`
	ReleaseDate          string               `json:"release_date"`
	ReleaseDatePrecision ReleaseDatePrecision `json:"release_date_precision"`
	DateStored           time.Time            `json:"-"`
}

type Albums []Album

func (a Albums) IDs() []string {
	var ids []string
	for _, album := range a {
		ids = append(ids, album.ID)
	}

	return ids
}
