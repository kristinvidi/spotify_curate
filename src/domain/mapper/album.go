package mapper

import (
	"time"

	db "src/db/model"
	api "src/spotifyapi/model"
)

func DBAAbumsAndArtistMappingFromGetArtistsAlbumsResponse(artistIDToAlbumsResponses map[api.ID][]*api.GetArtistsAlbumsResponse) (db.Albums, db.ArtistAlbumIDMappings) {
	var albums db.Albums
	var mappings db.ArtistAlbumIDMappings

	uniqueAlbums := make(map[api.ID]struct{})
	for artistID, responses := range artistIDToAlbumsResponses {
		for _, r := range responses {
			if r != nil {
				for _, a := range r.Albums {
					albums, mappings = addAPIAlbumToDBAlbumsAndMappings(artistID, a, uniqueAlbums, albums, mappings)
				}
			}
		}
	}

	return albums, mappings
}

func addAPIAlbumToDBAlbumsAndMappings(artistID api.ID, album api.Album, uniqueAlbums map[api.ID]struct{}, dbAlbums db.Albums, dbMappings db.ArtistAlbumIDMappings) (db.Albums, db.ArtistAlbumIDMappings) {
	if _, ok := uniqueAlbums[album.ID]; !ok {
		if album.AlbumType != api.AlbumTypeCompilation {
			uniqueAlbums[album.ID] = struct{}{}
			dbAlbum := dbAlbumFromAPIAlbum(album)
			dbAlbums = append(dbAlbums, dbAlbum)
			dbMapping := artistMappingFromAPIArtistAndAlbum(artistID, album)
			dbMappings = append(dbMappings, dbMapping)
		}
	}

	return dbAlbums, dbMappings
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
