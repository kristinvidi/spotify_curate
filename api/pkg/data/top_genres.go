package data

import (
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/model"
)

func addGenreCountToMap(genres apptype.Genres, mapGenreToGenreCount map[apptype.Genre]int) {
	for _, genre := range genres {
		if count, ok := mapGenreToGenreCount[genre]; ok {
			mapGenreToGenreCount[genre] = count + 1
		} else {
			mapGenreToGenreCount[genre] = 1
		}
	}
}

func GetMapOfGenreToCountOfGenre(artists []model.Artist) map[apptype.Genre]int {
	mapGenreToGenreCount := make(map[apptype.Genre]int)

	for _, artist := range artists {
		addGenreCountToMap(artist.Genres, mapGenreToGenreCount)
	}

	return mapGenreToGenreCount
}

func addArtistToMapOfGenre(artist model.Artist, mapGenreToGenreCount map[apptype.Genre][]string) {
	for _, genre := range artist.Genres {
		if artists, ok := mapGenreToGenreCount[genre]; ok {
			mapGenreToGenreCount[genre] = append(artists, artist.Name)
		} else {
			mapGenreToGenreCount[genre] = []string{artist.Name}
		}
	}
}

func GetMapOfGenreToArtists(artists []model.Artist) map[apptype.Genre][]string {
	mapGenreToArtistSlice := make(map[apptype.Genre][]string)

	for _, artist := range artists {
		addArtistToMapOfGenre(artist, mapGenreToArtistSlice)
	}

	return mapGenreToArtistSlice
}
