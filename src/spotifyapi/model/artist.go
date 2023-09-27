package model

type Artist struct {
	URI    URI       `json:"uri"`
	ID     SpotifyID `json:"id"`
	Name   string    `json:"name"`
	Genres []string  `json:"genres"`
}
