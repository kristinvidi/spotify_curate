package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"src/spotifyapi/constants"
	"src/spotifyapi/model"
)

type Playlist struct{}

func NewPlaylist() *Playlist {
	return &Playlist{}
}

func (p *Playlist) BuildCreatePlaylistRequest(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error) {
	if inputs.IDString() == nil {
		return nil, fmt.Errorf("cannot create a playlist without a user ID specified")
	}

	if inputs.CreatePlaylistInputs() == nil {
		return nil, fmt.Errorf("cannot create a playlist without playlist inputs specified")
	}

	path := filepath.Join(constants.URLPathUsers, *inputs.IDString(), constants.URLSpecifierPlaylists)

	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path,
	}
	baseURL := url.String()

	parameters := new(bytes.Buffer)

	json.NewEncoder(parameters).Encode(*inputs.CreatePlaylistInputs())

	req, err := http.NewRequest(http.MethodPost, baseURL, parameters)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func (p *Playlist) DecodeCreatePlaylistResponse(response http.Response) (*model.CreatePlaylistResponse, error) {
	var responseModel model.CreatePlaylistResponse
	err := json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (p *Playlist) BuildAddTracksToPlaylistRequest(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error) {
	if inputs.IDString() == nil {
		return nil, fmt.Errorf("cannot add tracks to a playlist without a playlist ID specified")
	}

	if inputs.URIs() == nil {
		return nil, fmt.Errorf("cannot add tracks to a playlist without track URIs specified")
	}

	path := filepath.Join(constants.URLPathPlaylists, *inputs.IDString(), constants.URLSpecifierTracks)

	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path,
	}
	baseURL := url.String()
	parameters := new(bytes.Buffer)
	json.NewEncoder(parameters).Encode(model.AddTracksToPlaylistInputs{TrackURIs: inputs.URIs()})

	req, err := http.NewRequest(http.MethodPost, baseURL, parameters)
	if err != nil {
		return nil, err
	}
	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func (p *Playlist) DecodeAddTracksToPlaylistResponse(response http.Response) (*model.AddTracksToPlaylistResponse, error) {
	var responseModel model.AddTracksToPlaylistResponse
	err := json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
