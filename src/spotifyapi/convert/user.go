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

func (u *User) BuildCurrentUsersProfileRequest(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error) {
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

func (u *User) BuildGetFollowedArtistsRequest(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   constants.URLPathMeFollowing,
	}
	url.RawQuery = u.encodeArtistTypeAndLimit(*inputs)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func (u *User) DecodeGetFollowedArtistsResponse(response http.Response) (*model.GetFollowedArtistsResponse, error) {
	var decodedResponse model.GetFollowedArtistsResponse
	err := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		return nil, err
	}

	return &decodedResponse, nil
}

func (u *User) encodeArtistTypeAndLimit(inputs model.RequestInput) string {
	url := url.URL{}
	params := url.Query()

	params.Set(constants.ParameterType, constants.TypeArtist)
	params.Set(constants.ParameterLimit, *inputs.BatchSize())

	if inputs.After() == nil {
		return params.Encode()
	}

	params.Set(constants.ParameterAfter, *inputs.After())

	return params.Encode()
}

func (u *User) BuildGetUsersSavedTracksRequest(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   constants.URLPathMeTracks,
	}
	url.RawQuery = u.encodeOffsetAndLimit(*inputs)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func (u *User) encodeOffsetAndLimit(inputs model.RequestInput) string {
	url := url.URL{}
	params := url.Query()

	params.Set(constants.ParameterLimit, *inputs.BatchSize())
	params.Set(constants.ParameterOffset, *inputs.Offset())

	if inputs.After() == nil {
		return params.Encode()
	}

	params.Set(constants.ParameterAfter, *inputs.After())

	return params.Encode()
}

func (u *User) DecodeGetUsersSavedTracksResponse(response http.Response) (*model.GetUsersSavedTracksResponse, error) {
	var decodedResponse model.GetUsersSavedTracksResponse
	err := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		return nil, err
	}

	return &decodedResponse, nil
}
