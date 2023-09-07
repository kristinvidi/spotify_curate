package client

import (
	"net/http"

	"src/spotifyapi/constants"
)

type Http struct {
	client *http.Client
}

func NewHttp() *Http {
	return &Http{client: &http.Client{}}
}

// DoRequest executes the Http request and returns the response. If the request fails due to an authentication error, the method will return ErrUnauthorized. If the response status code is not a successful code, the method will return ErrUnsuccessful. Otherwise, the response is returned with nil error.
func (h *Http) DoRequest(request *http.Request) (*http.Response, error) {
	response, err := h.client.Do(request)
	if err != nil {
		return nil, err
	}

	if !h.statusCodeIsAuthorized(response.StatusCode) {
		return nil, constants.ErrUnauthorized
	}

	if !h.statusCodeIsOK(response.StatusCode) {
		return nil, constants.ErrUnsuccessful
	}

	return response, nil
}

func (h *Http) statusCodeIsOK(code int) bool {
	return 200 <= code && code < 300
}

func (h *Http) statusCodeIsAuthorized(code int) bool {
	return code != http.StatusUnauthorized && code != http.StatusBadRequest
}
