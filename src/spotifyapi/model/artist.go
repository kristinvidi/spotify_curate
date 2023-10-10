package model

type Artist struct {
	URI    URI       `json:"uri"`
	ID     SpotifyID `json:"id"`
	Name   string    `json:"name"`
	Genres []string  `json:"genres"`
}

type Artists []Artist

type GetArtistsAlbumsResponse struct {
	HREF     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Albums   Albums `json:"items"`
}
