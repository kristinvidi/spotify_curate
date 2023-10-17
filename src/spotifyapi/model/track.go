package model

type Track struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
	// Popularity int    `json:"popularity"`
	Type string `json:"type"`
	URI  URI    `json:"uri"`
	// Album      Album   `json:"album"`
	// Artists    Artists `json:"artists"`
}

type Tracks []Track

type SimplifiedTrack struct {
	Name string `json:"name"`
	ID   ID     `json:"id"`
	URI  URI    `json:"uri"`
}
