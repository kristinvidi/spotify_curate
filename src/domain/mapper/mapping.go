package mapper

import (
	db "src/db/model"
	"src/domain/model"
)

func DBUserArtistIDGenreMappingsFromDBGenreAndArtists(genre db.UserIDGenreMapping, artists model.Artists) db.UserArtistIDGenreMappings {
	mappings := make(db.UserArtistIDGenreMappings, len(artists))
	for i, artist := range artists {
		mappings[i] = db.UserArtistIDGenreMapping{
			UserID:   genre.UserID,
			ArtistID: IDToDBID(artist.ID),
			GenreID:  genre.ID,
		}
	}

	return mappings
}
