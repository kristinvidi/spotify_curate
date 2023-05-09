package httprequest

import (
	"fmt"
	"net/http"

	"github.com/openlyinc/pointy"
)

type HttpRequest struct {
	httpClient *http.Client
}

func NewHttpRequest(httpClient *http.Client) HttpRequest {
	return HttpRequest{
		httpClient: httpClient,
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

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("http request failed with status: %s", response.Status)
	}

	return nil
}

func (h *HttpRequest) CheckResponseStatusIsBadRequest(request *http.Request) (*bool, error) {
	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return checkResponseStatusIsBadRequest(response)
}

func checkResponseStatusIsBadRequest(response *http.Response) (*bool, error) {
	if response == nil {
		return nil, fmt.Errorf("cannot infer status for nil response")
	}

	if response.StatusCode != http.StatusBadRequest {
		return pointy.Bool(false), nil
	}

	return pointy.Bool(true), nil
}
