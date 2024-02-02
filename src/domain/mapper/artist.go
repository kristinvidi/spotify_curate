package mapper

import (
	db "src/db/model"
	"src/domain/model"
)

func ArtistsFromDBArtists(dbArtists db.Artists) model.Artists {
	artists := make(model.Artists, len(dbArtists))
	for i, a := range dbArtists {
		artists[i] = ArtistFromDBArtist(a)
	}

	return artists
}

func ArtistFromDBArtist(dbArtist db.Artist) model.Artist {
	return model.Artist{
		ID:        model.ID(dbArtist.ID),
		URI:       model.URI(dbArtist.URI),
		Name:      dbArtist.Name,
		CreatedAt: dbArtist.CreatedAt,
	}
}
