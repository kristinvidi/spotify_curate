package query

import (
	"src/db/constants"
	"src/db/model"
)

func (p *PostgresDB) InsertUserData(user model.User) error {
	return p.insert(&user, constants.ColumnSpotifyID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertUserToArtistSpotifyIDMappings(mappings []model.UserArtistSpotifyIDMapping) error {
	return p.insert(&mappings, constants.ColumnUserArtistSpotifyID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertAlbums(albums model.Albums) error {
	return p.insert(&albums, constants.ColumnSpotifyID, constants.OnConflictDoNothing)
}
