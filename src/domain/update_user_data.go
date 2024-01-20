package domain

import (
	"src/config"
	"src/db/query"
	"src/domain/mapper"
	"src/domain/model"
	"src/spotifyapi"
	"src/spotifyapi/api"

	"go.uber.org/zap"
)

type UserUpdater struct {
	config  *config.Config
	userAPI *api.User
	db      *query.PostgresDB
	logger  *zap.Logger
}

func NewUserUpdater(config *config.Config, logger *zap.Logger) *UserUpdater {
	return &UserUpdater{
		config:  config,
		userAPI: spotifyapi.GetUser(config),
		db:      query.NewPostgresDB(config.Database),
		logger:  logger,
	}
}

func (u *UserUpdater) UpdateUserData() error {
	user, err := u.getAndStoreCurrentUserProfile()
	if err != nil {
		return err
	}

	responses, err := u.userAPI.GetCurrentUsersFollowedArtists()
	if err != nil {
		return err
	}

	// Insert artist data
	dbArtists := mapper.DBFollowedArtistsFromGetFollowedArtistsResponse(responses)

	found := 0
	for i, a := range dbArtists {
		if a.ID == "Barkhan" {
			found = i
		}
	}

	u.logger.Info("found artist?", zap.Int("position", found))

	err = u.db.InsertArtistData(dbArtists)
	if err != nil {
		return err
	}

	// Remove unfollowed artists genre mappings
	rowsDeleted, err := u.db.DeleteUserToArtistIDGenreMappings(mapper.IDToDBID(user.ID), dbArtists.IDs())
	if err != nil {
		return err
	}

	u.logger.Info("successfully removed user artist genre mappings", zap.Int64("rowsDeleted", rowsDeleted))

	// Remove unfollowed artists mappings
	rowsDeleted, err = u.db.DeleteUserToArtistIDMappings(mapper.IDToDBID(user.ID), dbArtists.IDs())
	if err != nil {
		return err
	}

	u.logger.Info("successfully removed user artist mappings", zap.Int64("rowsDeleted", rowsDeleted))

	// Insert user to artist mapping data
	err = u.db.InsertUserToArtistIDMappings(
		mapper.DBUserToArtistMappingFromGetFollowedArtistsResponse(user.ID, responses),
	)
	if err != nil {
		return err
	}

	artistIDToAlbumsResponses, err := u.userAPI.GetCurrentUsersFollowedArtistsToAlbums(
		mapper.APIArtistsFromGetFollowedArtistsResponse(responses),
	)
	if err != nil {
		return err
	}

	albums, artistToAlbumMappings := mapper.DBAAbumsAndArtistMappingFromGetArtistsAlbumsResponse(artistIDToAlbumsResponses)

	err = u.db.InsertAlbums(albums)
	if err != nil {
		return err
	}

	err = u.db.InsertArtistAlbumIDMappings(artistToAlbumMappings)
	if err != nil {
		return err
	}

	err = u.db.InsertUserUpdateStatus(
		mapper.UserUpdateStatus(user.ID),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUpdater) GetUnmappedArtistsForUser() ([]model.Artist, error) {
	user, err := u.getAndStoreCurrentUserProfile()
	if err != nil {
		return nil, err
	}

	artists, err := u.db.GetUnmappedArtistsForUser(mapper.IDToDBID(user.ID))
	if err != nil {
		return nil, err
	}

	return mapper.ArtistsFromDBArtists(artists), nil
}

func (u *UserUpdater) getAndStoreCurrentUserProfile() (*model.User, error) {
	response, err := u.userAPI.GetCurrentUsersProfile()
	if err != nil {
		return nil, err
	}

	user := mapper.UserFromCurrentUsersProfileResponse(response)

	err = u.db.InsertUserData(
		*mapper.UserToDBUser(user),
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
