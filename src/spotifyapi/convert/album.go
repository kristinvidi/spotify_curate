package convert

import (
	"net/http"
	"net/url"
	"path"
	"src/spotifyapi/constants"
	"src/spotifyapi/model"
)

type Album struct{}

func NewAlbum() *Album {
	return &Album{}
}

func (a *Album) BuildGetAlbumTracksRequest(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path.Join(constants.URLPathAlbums, *inputs.IDString(), constants.TypeTracks),
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}
