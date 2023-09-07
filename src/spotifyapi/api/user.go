package api

import (
	"src/spotifyapi/convert"
	"src/spotifyapi/model"
)

type User struct {
	api       *API
	converter *convert.User
}

func NewUser(api *API, converter *convert.User) *User {
	return &User{api: api}
}

func (u *User) GetCurrentUsersProfile() (*model.GetCurrentUsersProfileResponse, error) {
	response, err := u.api.DoRequest(u.converter.BuildCurrentUsersProfileRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decodedResponse, err := u.converter.DecodeCurrentUsersProfileResponse(*response)
	if err != nil {
		return nil, err
	}

	return decodedResponse, nil
}
