package converter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path/filepath"
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/constants"
	"spotify_app/api/pkg/model"
)

func BuildCreatePlaylistRequest(accessToken apptype.AccessToken, userID apptype.UserID, input model.CreatePlaylistRequest) (*http.Request, error) {
	path := filepath.Join(constants.URLPathUsers, string(userID), constants.URLSpecifierPlaylists)

	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path,
	}
	baseURL := url.String()

	parameters := new(bytes.Buffer)
	json.NewEncoder(parameters).Encode(input)

	req, err := http.NewRequest(http.MethodPost, baseURL, parameters)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func DecodeCreatePlaylistResponse(response http.Response) (*model.CreatePlaylistResponse, error) {
	var responseModel model.CreatePlaylistResponse
	err := json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func BuildAddTracksToPlaylistRequests(accessToken apptype.AccessToken, playlistID apptype.PlaylistID, tracks model.Tracks) (*[]http.Request, error) {
	path := filepath.Join(constants.URLPathPlaylists, string(playlistID), constants.URLSpecifierTracks)

	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path,
	}
	baseURL := url.String()

	var requests []http.Request
	batchSize := 100

	for i := 0; i < len(tracks); i += batchSize {
		t := getSubsetOfItem(tracks, i, batchSize)
		r, err := buildAddTracksToPlaylistRequests(t, baseURL, accessToken)
		if err != nil {
			return nil, err
		}
		requests = append(requests, *r)
	}

	return &requests, nil
}

func buildAddTracksToPlaylistRequests(tracks model.Tracks, url string, accessToken apptype.AccessToken) (*http.Request, error) {
	input := model.NewAddItemsToPlaylistRequest(tracks.URIs())
	parameters := new(bytes.Buffer)
	json.NewEncoder(parameters).Encode(input)

	req, err := http.NewRequest(http.MethodPost, url, parameters)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func getSubsetOfItem[V any](s []V, start, batchSize int) []V {
	end := start + batchSize
	if end > len(s) {
		end = len(s)
	}

	return s[start:end]
}
