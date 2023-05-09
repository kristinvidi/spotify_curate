package converter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/constants"
	"spotify_app/api/pkg/model"
)

func BuildGetCurrentUsersProfileRequest(accessToken apptype.AccessToken) (*http.Request, error) {
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

func DecodeGetCurrentUsersProfileResponse(response http.Response) (*model.GetCurrentUsersProfileResponse, error) {
	var responseModel model.GetCurrentUsersProfileResponse
	err := json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func BuildGetFollowedArtistsRequest(accessToken apptype.AccessToken, after *string) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   constants.URLPathFollowing,
	}
	url.RawQuery = encodedQueryParameters(after, 50)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func DecodeGetFollowedArtistsResponse(response http.Response) (*model.GetFollowedArtistsResponse, error) {
	var decodedResponse model.GetFollowedArtistsResponse
	err := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		return nil, err
	}

	return &decodedResponse, nil
}

func encodedQueryParameters(after *string, limit int) string {
	url := url.URL{}
	params := url.Query()
	params.Set(constants.ParameterType, constants.TypeArtist)
	params.Set(constants.ParameterLimit, fmt.Sprint(limit))

	if after == nil {
		return params.Encode()
	}

	params.Set(constants.ParameterAfter, *after)

	return params.Encode()
}
