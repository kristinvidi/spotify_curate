package mapper

import (
	db "src/db/model"
	"src/domain/model"
	api "src/spotifyapi/model"
)

func DBUserFromCurrentUsersProfileResponse(response *api.GetCurrentUsersProfileResponse) *db.User {
	user := userFromCurrentUsersProfileResponse(response)
	return userToDBUser(user)
}

func userFromCurrentUsersProfileResponse(response *api.GetCurrentUsersProfileResponse) *model.User {
	if response == nil {
		return nil
	}

	return &model.User{
		DisplayName: model.Name(response.DisplayName),
		Email:       model.Email(response.Email),
		ID:          model.ID(response.ID),
		URI:         model.URI(response.URI),
		Country:     model.CountryCode(response.Country),
	}
}

func userToDBUser(user *model.User) *db.User {
	if user == nil {
		return nil
	}

	return &db.User{
		DisplayName: string(user.DisplayName),
		Email:       string(user.Email),
		ID:          string(user.ID),
		URI:         string(user.URI),
		Country:     string(user.Country),
	}
}
