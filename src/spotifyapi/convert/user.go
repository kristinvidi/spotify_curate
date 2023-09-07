package convert

import (
	"encoding/json"
	"net/http"
	"net/url"

	"src/spotifyapi/constants"
	"src/spotifyapi/model"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) BuildCurrentUsersProfileRequest(accessToken model.AccessToken) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   constants.URLPathMe,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func (u *User) DecodeCurrentUsersProfileResponse(response http.Response) (*model.GetCurrentUsersProfileResponse, error) {
	var responseModel model.GetCurrentUsersProfileResponse
	err := json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
