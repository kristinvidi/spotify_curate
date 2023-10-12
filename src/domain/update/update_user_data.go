package update

import (
	"src/config"
	"src/db/query"
	"src/domain/mapper"
	"src/domain/model"
	"src/spotifyapi"
	"src/spotifyapi/api"
)

type UserData struct {
	config  *config.Config
	userAPI *api.User
	db      *query.PostgresDB
}

func NewUserData(config *config.Config) *UserData {
	return &UserData{
		config:  config,
		userAPI: spotifyapi.GetUser(config),
		db:      query.NewPostgresDB(config.Database),
	}
}

func (u *UserData) UpdateAllUserData() error {
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

	allAlbums, err := u.userAPI.GetCurrentUsersFollowedArtistsAlbums(
		mapper.APIArtistsFromGetFollowedArtistsResponse(responses),
	)
	if err != nil {
		return err
	}

	err = u.db.InsertAlbums(
		mapper.DBAAbumsFromGetArtistsAlbumsResponse(allAlbums),
	)
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

func (u *UserData) getAndStoreCurrentUserProfile() (*model.User, error) {
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
