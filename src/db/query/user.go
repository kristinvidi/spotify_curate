package query

import (
	"src/db/constants"
	"src/db/model"
)

func (p *PostgresDB) InsertUserData(user model.User) error {
	return p.insertWithConflict(&user, constants.ColumnSpotifyID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertUserToArtistSpotifyIDMappings(mappings []model.UserArtistSpotifyIDMapping) error {
	return p.insertWithConflict(&mappings, constants.ColumnUserArtistSpotifyID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertAlbums(albums model.Albums) error {
	return p.insertWithConflict(&albums, constants.ColumnSpotifyID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertUserUpdateStatus(status model.UserUpdateStatus) error {
	return p.insertNoConflict(&status)
}
