package converter

import (
	"encoding/json"
	"net/http"
	"net/url"

	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/constants"
	"spotify_app/api/pkg/model"
)

func BuildGetFollowedArtistsRequest(accessToken apptype.AccessToken, after *string) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   constants.URLPathFollowing,
	}
	url.RawQuery = encodedQueryParameters(after)

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

func encodedQueryParameters(after *string) string {
	url := url.URL{}
	params := url.Query()
	params.Set(constants.ParameterType, constants.TypeArtist)
	params.Set(constants.ParameterLimit, "50")

	if after == nil {
		return params.Encode()
	}

	params.Set(constants.ParameterAfter, *after)

	return params.Encode()
}
