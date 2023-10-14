package domain

import (
	"src/config"
	"src/db/query"
	"src/domain/mapper"
	"src/domain/model"
	"src/spotifyapi"
	"src/spotifyapi/api"
)

type UserUpdater struct {
	config  *config.Config
	userAPI *api.User
	db      *query.PostgresDB
}

func NewUserUpdater(config *config.Config) *UserUpdater {
	return &UserUpdater{
		config:  config,
		userAPI: spotifyapi.GetUser(config),
		db:      query.NewPostgresDB(config.Database),
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
	err = u.db.InsertArtistData(
		mapper.DBFollowedArtistsFromGetFollowedArtistsResponse(responses),
	)
	if err != nil {
		return err
	}

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
