package data

// import (
// 	apptype "spotify_app/api/pkg/app_type"
// 	"spotify_app/api/pkg/model"
// )

// func GetMapOfArtistToGenre(artists []model.Artist) (mapGenreToArtistSlice map[apptype.Genre]model.Artists, artistsWithoutGenres []model.Artist) {
// 	mapOfArtistToGenre = make(map[apptype.Genre]model.Artists)
// 	for _, artist := range artists {
// 		if len(artist.Genres) == 0 {
// 			artistsWithoutGenres = append(artistsWithoutGenres, artist)
// 		} else {
// 			addArtistToMapOfGenre(artist, mapGenreToArtistSlice)
// 		}
// 	}

// 	return mapGenreToArtistSlice, artistsWithoutGenres
// }

// func addArtistToMapOfGenre(artist model.Artist, mapGenreToGenreCount map[apptype.Genre]model.Artists) {
// 	genre := artist.Genres[0]
// 	if artists, ok := mapGenreToGenreCount[genre]; ok {
// 		mapGenreToGenreCount[genre] = append(artists, artist)
// 	} else {
// 		mapGenreToGenreCount[genre] = model.Artists{artist}
// 	}
// }
