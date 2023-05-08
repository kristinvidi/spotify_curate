package httprequest

import (
	"fmt"
	"net/http"
)

type HttpRequest struct {
	httpClient *http.Client
}

func NewHttpRequest(httpClient *http.Client) HttpRequest {
	return HttpRequest{
		httpClient: httpClient,
	}
}

func (h *HttpRequest) DoRequestAndCheckStatus(request *http.Request) (*http.Response, error) {
	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	err = checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func checkResponseStatus(response *http.Response) error {
	if response == nil {
		return fmt.Errorf("cannot infer status for nil response")
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("http request failed with status: %s", response.Status)
	}

	return nil
}
