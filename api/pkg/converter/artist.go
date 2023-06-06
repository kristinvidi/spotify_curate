package converter

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path/filepath"

	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/constants"
	"spotify_app/api/pkg/model"
)

func BuildGetArtistTopTracksRequest(accessToken apptype.AccessToken, artistID string, countryCode apptype.UserCountryCode) (*http.Request, error) {
	path := filepath.Join(constants.URLPathArtist, artistID, constants.URLSpecifierTopTracks)

	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path,
	}
	params := url.Query()
	params.Set(constants.ParameterMarket, string(countryCode))
	url.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func DecodeGetArtistTopTracksResponse(response http.Response) (*model.GetArtistTopTracksResponse, error) {
	var responseModel model.GetArtistTopTracksResponse
	err := json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func BuildGetArtistRelatedArtistsRequest(accessToken apptype.AccessToken, artistID string) (*http.Request, error) {
	path := filepath.Join(constants.URLPathArtist, artistID, constants.URLSpecifierRelatedArtists)

	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path,
	}
	params := url.Query()
	url.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func DecodeGetArtistRelatedArtistsResponse(response http.Response) (*model.GetArtistRelatedArtistsResponse, error) {
	var responseModel model.GetArtistRelatedArtistsResponse
	err := json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

// func BuildGetArtistsAlbumsRequest(accessToken apptype.AccessToken, artistID string, after *string, countryCode apptype.UserCountryCode) (*http.Request, error) {
// 	path := filepath.Join(constants.URLPathArtist, artistID, constants.URLSpecifierAlbums)

// 	url := url.URL{
// 		Scheme: constants.URLScheme,
// 		Host:   constants.URLHostAPI,
// 		Path:   path,
// 	}
// 	url.RawQuery = encodedQueryParameters(after, constants.ReplyLimit, &countryCode)

// 	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

// 	return req, nil
// }

// func DecodeGetArtistsAlbumsResponse(response http.Response) (*model.GetArtistsAlbumsResponse, error) {
// 	var decodedResponse model.GetArtistsAlbumsResponse
// 	err := json.NewDecoder(response.Body).Decode(&decodedResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &decodedResponse, nil
// }
