package model

import (
	"time"
)

type Album struct {
	ID        ID        `json:"id"`
	URI       URI       `json:"uri"`
	Name      string    `json:"name"`
	AlbumType AlbumType `json:"album_type"`
	Genres    []string  `json:"genres"`
	Artists   Artists   `json:"artists"`
	TrackList struct {
		Tracks Tracks `json:"items"`
	} `json:"tracks"`
	Offset               int
	TotalTracks          int                  `json:"total"`
	ReleaseDateString    string               `json:"release_date"`
	ReleaseDatePrecision ReleaseDatePrecision `json:"release_date_precision"`
}

func (a Album) ReleaseDate() time.Time {
	t, _ := time.Parse(string(YearMonthDay), a.ReleaseDateString)
	return t
}

type Albums []Album

type GetAlbumTracksResponse struct {
	Limit  int               `json:"limit"`
	Total  int               `json:"total"`
	Tracks []SimplifiedTrack `json:"items"`
}
