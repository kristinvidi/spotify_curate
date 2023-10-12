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
			albums = append(albums, dbAlbumsFromAPIAlbums(r.Albums)...)
		}
	}

	return albums
}

func dbAlbumsFromAPIAlbums(albums api.Albums) db.Albums {
	dbAlbums := make(db.Albums, len(albums))

	for i, a := range albums {
		dbAlbums[i] = dbAlbumFromAPIAlbum(a)
	}

	return dbAlbums
}

func dbAlbumFromAPIAlbum(album api.Album) db.Album {
	return db.Album{
		ID:                   db.ID(album.ID),
		URI:                  db.URI(album.URI),
		Name:                 album.Name,
		ReleaseDate:          album.ReleaseDate(),
		ReleaseDatePrecision: db.ReleaseDatePrecision(album.ReleaseDatePrecision),
		CreatedAt:            time.Now(),
	}
}
