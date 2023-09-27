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
	user, err := u.getCurrentUserProfile()
	if err != nil {
		return err
	}

	responses, err := u.userAPI.GetCurrentUsersFollowedArtists()
	if err != nil {
		return err
	}

	// Insert user data
	mappedUser := mapper.UserToDBUser(user)
	err = u.db.InsertUserData(*mappedUser)
	if err != nil {
		return err
	}

	// Insert artist data
	mappedArtists := mapper.DBFollowedArtistsFromGetFollowedArtistsResponse(responses)
	err = u.db.InsertArtistData(mappedArtists)
	if err != nil {
		return err
	}

	// Insert user to artist mapping data
	userToArtistMapping := mapper.DBUserToArtistMappingFromGetFollowedArtistsResponse(user.ID, responses)
	err = u.db.InsertUserToArtistSpotifyIDMappings(userToArtistMapping)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserData) getCurrentUserProfile() (*model.User, error) {
	response, err := u.userAPI.GetCurrentUsersProfile()
	if err != nil {
		return nil, err
	}

	return mapper.UserFromCurrentUsersProfileResponse(response), nil
}
