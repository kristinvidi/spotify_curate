package converter

// import (
// 	"net/http"
// 	"net/url"
// 	"path/filepath"

// 	apptype "spotify_app/api/pkg/app_type"
// 	"spotify_app/api/pkg/constants"
// )
//
// func BuildGetTrackRequest(accessToken apptype.AccessToken, trackID string) (*http.Request, error) {
// 	path := filepath.Join(constants.URLPathTrack, trackID)

// 	url := url.URL{
// 		Scheme: constants.URLScheme,
// 		Host:   constants.URLHostAPI,
// 		Path:   path,
// 	}

// 	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Add(constants.HeaderAuthorization, accessToken.HeaderValue())

// 	return req, nil
// }
//
// func DecodeGetTrackResponse(response http.Response) (*model.GetTrackResponse, error) {
// 	var responseModel model.GetTrackResponse
// 	err := json.NewDecoder(response.Body).Decode(&responseModel)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &responseModel, nil
// }
