package mapper

import (
	db "src/db/model"
	"time"

	api "src/spotifyapi/model"
)

func DBAAbumsAndArtistMappingFromGetArtistsAlbumsResponse(artistIDToAlbumsResponses map[api.ID][]*api.GetArtistsAlbumsResponse) (db.Albums, db.ArtistAlbumIDMappings) {
	var albums db.Albums
	var mappings db.ArtistAlbumIDMappings

	for id, responses := range artistIDToAlbumsResponses {
		for _, r := range responses {
			if r != nil {
				for _, a := range r.Albums {
					album := dbAlbumFromAPIAlbum(a)
					mapping := artistMappingFromAPIArtistAndAlbum(id, a)
					albums = append(albums, album)
					mappings = append(mappings, mapping)
				}
			}
		}
	}

	return albums, mappings
}

func dbAlbumFromAPIAlbum(album api.Album) db.Album {
	return db.Album{
		ID:                   db.ID(album.ID),
		URI:                  db.URI(album.URI),
		Name:                 album.Name,
		Type:                 db.AlbumType(album.AlbumType),
		ReleaseDate:          album.ReleaseDate(),
		ReleaseDatePrecision: db.ReleaseDatePrecision(album.ReleaseDatePrecision),
		CreatedAt:            time.Now(),
	}
}

func artistMappingFromAPIArtistAndAlbum(artistID api.ID, album api.Album) db.ArtistAlbumIDMapping {
	return db.ArtistAlbumIDMapping{
		ArtistID: db.ID(artistID),
		AlbumID:  db.ID(album.ID),
	}
}
