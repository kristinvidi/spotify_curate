package model

import (
	// "encoding/json"
	"encoding/json"
	apptype "spotify_app/api/pkg/app_type"
)

type Artist struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Genres apptype.Genres `json:"genres"`
	// KikiLabel *apptype.Genre
}

func (a *Artist) SetGenres(g apptype.Genres) {
	a.Genres = g
}

func (a *Artist) IDAndGenreKey() string {
	return a.ID + "#" + string(a.Genres[0])
}

func (a *Artist) ToJSON() ([]byte, error) {
	jsonArtist, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return jsonArtist, nil
}

type Artists []Artist

// func (a *Artists) ToJSON() ([]byte, error) {
// 	jsonArtist, err := json.Marshal(a)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return jsonArtist, nil
// }

func (a Artists) ArtistIDToJSONArtist() ([]byte, error) {
	idToArtistJSON := make(map[string]Artist)
	for _, artist := range a {
		idToArtistJSON[artist.ID] = artist
	}

	artistsJSON, err := json.Marshal(idToArtistJSON)
	if err != nil {
		return nil, err
	}

	return artistsJSON, nil
}

type GetArtistTopTracksResponse struct {
	Tracks Tracks `json:"tracks"`
}

type GetArtistRelatedArtistsResponse struct {
	Artists Artists `json:"artists"`
}

type GetArtistsAlbumsResponse struct {
	HREF     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Albums   Albums `json:"items"`
}

func (a Artists) TopGenre() *apptype.Genre {
	genreToCount := make(map[apptype.Genre]int)

	for _, artist := range a {
		a.addGenresToMap(artist.Genres, genreToCount)
	}

	return a.getTopGenreFromMapOfGenreToCount(genreToCount)
}

func (a Artists) addGenresToMap(genres apptype.Genres, genreToCount map[apptype.Genre]int) {
	for _, genre := range genres {
		genreToCount[genre] += 1
	}
}

func (a Artists) getTopGenreFromMapOfGenreToCount(genreToCount map[apptype.Genre]int) *apptype.Genre {
	var topGenre apptype.Genre
	topCount := 0
	for genre, count := range genreToCount {
		topGenre, topCount = a.reassignTopGenreIfCountIsGreaterThanPreviousTopCount(genre, count, topGenre, topCount)
	}

	if topGenre == "" {
		return nil
	}

	return ifValueIsEmptyStringReturnNilElseReturnString(topGenre)
}

func (a Artists) reassignTopGenreIfCountIsGreaterThanPreviousTopCount(genre apptype.Genre, genreCount int, currentTopGenre apptype.Genre, currentTopCount int) (apptype.Genre, int) {
	if genreCount <= currentTopCount {
		return currentTopGenre, currentTopCount
	}

	return genre, genreCount
}

func ifValueIsEmptyStringReturnNilElseReturnString(value apptype.Genre) *apptype.Genre {
	if value == "" {
		return nil
	}

	return &value
}
