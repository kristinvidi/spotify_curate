package model

import "time"

type Album struct {
	ID        int64     `bun:",pk,autoincrement"`
	SpotifyID int64     `bun:",unique"`
	Name      string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,default:now()"`
	// AlbumType string         `json:"album_type"`
	// Genres    apptype.Genres `json:"genres"`
	// Artists   Artists        `json:"artists"`
	// TrackList struct {
	// 	Tracks Tracks `json:"items"`
	// } `json:"tracks"`
	// TotalTracks          int                  `json:"total_tracks"`
	// ReleaseDate          string               `json:"release_date"`
	// ReleaseDatePrecision ReleaseDatePrecision `json:"release_date_precision"`
	// DateStored           time.Time            `json:"-"`
}
