package query

import (
	"src/db/constants"
	"src/db/model"
)

func (p *PostgresDB) InsertUserData(user model.User) error {
	return p.insertWithConflict(&user, constants.ColumnID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertUserToArtistIDMappings(mappings []model.UserArtistIDMapping) error {
	return p.insertWithConflict(&mappings, constants.ColumnUserArtistID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) DeleteUserToArtistIDGenreMappings(userID model.ID, artistIDs []model.ID) (int64, error) {
	return p.deleteByUserInAndArtistIDNotIn((*model.UserArtistIDGenreMapping)(nil), userID, artistIDs)
}

func (p *PostgresDB) DeleteUserToArtistIDMappings(userID model.ID, artistIDs []model.ID) (int64, error) {
	return p.deleteByUserInAndArtistIDNotIn((*model.UserArtistIDMapping)(nil), userID, artistIDs)
}

func (p *PostgresDB) InsertAlbums(albums model.Albums) error {
	return p.insertWithConflict(&albums, constants.ColumnID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertArtistAlbumIDMappings(mappings model.ArtistAlbumIDMappings) error {
	return p.insertWithConflict(&mappings, constants.ColumnArtistAlbumID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) InsertUserUpdateStatus(status model.UserUpdateStatus) error {
	return p.insertNoConflict(&status)
}
