package httprequest

import (
	"fmt"
	"net/http"

	"github.com/openlyinc/pointy"

	"spotify_app/api/config"
)

type HttpRequest struct {
	httpClient    *http.Client
	configManager *config.Manager
}

func NewHttpRequest(httpClient *http.Client, configManager *config.Manager) HttpRequest {
	return HttpRequest{
		httpClient:    httpClient,
		configManager: configManager,
	}
}

func (h *HttpRequest) DoRequestAndCheckStatusIsOK(request *http.Request) (*http.Response, error) {
	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	err = checkResponseStatusIsOK(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func checkResponseStatusIsOK(response *http.Response) error {
	if response == nil {
		return fmt.Errorf("cannot infer status for nil response")
	}

	if !statusCodeIsOK(response.StatusCode) {
		return fmt.Errorf("http request failed with status: %s", response.Status)
	}

	return nil
}

func statusCodeIsOK(code int) bool {
	return 200 <= code && code < 300
}

func (h *HttpRequest) checkResponseStatusIsBadRequest(request *http.Request) (*bool, error) {
	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return checkResponseStatusIsUnauthorized(response)
}

func checkResponseStatusIsUnauthorized(response *http.Response) (*bool, error) {
	if response == nil {
		return nil, fmt.Errorf("cannot infer status for nil response")
	}

	if response.StatusCode != http.StatusUnauthorized && response.StatusCode != http.StatusBadRequest {
		return pointy.Bool(false), nil
	}

	return pointy.Bool(true), nil
}
