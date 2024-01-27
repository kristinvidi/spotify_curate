package convert

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"src/spotifyapi/constants"
	"src/spotifyapi/model"
)

type Artist struct{}

func NewArtist() *Artist {
	return &Artist{}
}

func (a *Artist) BuildGetArtistsAlbumsRequest(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAPI,
		Path:   path.Join(constants.URLPathArtist, *inputs.IDString(), constants.TypeAlbums),
	}
	url.RawQuery = a.encodeLimitAndOffset(*inputs)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

	return req, nil
}

func (a *Artist) DecodeGetArtistsAlbumsResponse(response http.Response) (*model.GetArtistsAlbumsResponse, error) {
	var decodedResponse model.GetArtistsAlbumsResponse
	err := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		return nil, err
	}

	return &decodedResponse, nil
}

func (a *Artist) encodeLimitAndOffset(inputs model.RequestInput) string {
	url := url.URL{}
	params := url.Query()

	params.Set(constants.ParameterLimit, *inputs.BatchSize())
	params.Set(constants.ParameterOffset, *inputs.Offset())

	return params.Encode()
}
