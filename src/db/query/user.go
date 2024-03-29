package query

import (
	"context"

	"src/db/constants"
	"src/db/model"

	"github.com/uptrace/bun"
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

func (p *PostgresDB) GetArtistIDsWithMappingsForUser(userID model.ID) ([]model.ID, error) {
	var artistIDs []model.ID

	err := p.db.NewSelect().
		ColumnExpr(string(constants.ColumnArtistID)).
		Model(&artistIDs).
		Table("user_artist_spotify_id_genre_mapping").
		Where("? = ?", bun.Ident(constants.ColumnUserID), userID).
		Group(string(constants.ColumnArtistID)).
		Scan(context.Background())

	return artistIDs, err
}

func (p *PostgresDB) GetUnmappedArtistsForUser(userID model.ID) (model.Artists, error) {
	mappedArtistsSubquery := p.db.NewSelect().
		ColumnExpr(constants.ColumnArtistID.String()).
		Table("user_artist_spotify_id_genre_mapping").
		Where("? = ?", bun.Ident(constants.ColumnUserID), userID).
		Group(constants.ColumnArtistID.String())

	unmappedArtistsSubquery := p.db.NewSelect().
		ColumnExpr(constants.ColumnArtistID.String()).
		Table("user_artist_spotify_id_mapping").
		Where("? = ?", bun.Ident(constants.ColumnUserID), userID).
		Where("? not in (?)", bun.Ident(constants.ColumnArtistID), mappedArtistsSubquery)

	var artists model.Artists
	err := p.db.NewSelect().
		Model(&artists).
		Where("? in (?)", bun.Ident(constants.ColumnID), unmappedArtistsSubquery).Scan(context.Background())

	return artists, err
}

func (p *PostgresDB) GetMappedArtistsForUserByArtistNames(userID model.ID, artistNames []string) (model.Artists, error) {
	var artists model.Artists

	err := p.db.NewSelect().
		Model(&artists).
		Join("INNER JOIN user_artist_spotify_id_mapping uasim ON uasim.artist_spotify_id = artist.spotify_id").
		Where("? = ?", bun.Ident(constants.ColumnUserID), userID).
		Where("? in (?)", bun.Ident(constants.ColumnDisplayName), bun.In(artistNames)).
		Scan(context.Background())

	return artists, err
}

func (p *PostgresDB) InsertUserPlaylistTrackIDMappings(mappings []model.UserPlaylistTrackIDMapping) error {
	return p.insertWithConflict(&mappings, constants.ColumnUserPlaylistTrackID, constants.OnConflictDoNothing)
}

func (p *PostgresDB) UpsertUserSavedTracks(userID model.ID, userSavedTracks []model.UserSavedTracks) error {
	return p.db.RunInTx(context.Background(), nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := p.deleteByUserID(&[]model.UserSavedTracks{}, userID)
		if err != nil {
			return err
		}

		return p.insertNoConflict(&userSavedTracks)
	})
}
