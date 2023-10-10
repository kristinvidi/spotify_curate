package mapper

import (
	db "src/db/model"
	"time"

	api "src/spotifyapi/model"
)

func DBAAbumsFromGetArtistsAlbumsResponse(responses []*api.GetArtistsAlbumsResponse) db.Albums {
	var albums db.Albums
	for _, r := range responses {
		if r != nil {
			albums = append(albums, DBAlbumsFromAPIAlbums(r.Albums)...)
		}
	}

	return albums
}

func DBAlbumsFromAPIAlbums(albums api.Albums) db.Albums {
	dbAlbums := make(db.Albums, len(albums))

	for i, a := range albums {
		album := db.Album{
			SpotifyID:            a.SpotifyID,
			URI:                  a.URI,
			Name:                 a.Name,
			ReleaseDate:          a.ReleaseDate(),
			ReleaseDatePrecision: db.ReleaseDatePrecision(a.ReleaseDatePrecision),
			CreatedAt:            time.Now(),
		}

		dbAlbums[i] = album
	}

	return dbAlbums
}
