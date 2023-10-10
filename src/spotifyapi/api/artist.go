package api

import "src/spotifyapi/convert"

type Artist struct {
	api       *API
	converter *convert.Artist
}

func NewArtist(api *API, converter *convert.Artist) *Artist {
	return &Artist{api: api, converter: converter}
}
