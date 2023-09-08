package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/pkg/browser"

	"src/config"
	"src/spotifyapi/authentication"
	"src/spotifyapi/client"
	"src/spotifyapi/constants"
	"src/spotifyapi/convert"
	"src/spotifyapi/model"
)

const delay = time.Second

type API struct {
	httpClient   *client.Http
	config       *config.Config
	tokenStorage *authentication.AccessTokenStorage
}

func NewAPI(httpClient *client.Http, config *config.Config, tokenStorage *authentication.AccessTokenStorage) *API {
	return &API{httpClient: httpClient, config: config, tokenStorage: tokenStorage}
}

func (a *API) DoRequest(requestBuilder func(accessToken model.AccessToken) (*http.Request, error)) (*http.Response, error) {
	if a.config == nil {
		return nil, constants.ErrMissingConfig
	}

	resp, err := a.executeRequest(requestBuilder)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *API) executeRequest(requestBuilder func(accessToken model.AccessToken) (*http.Request, error)) (*http.Response, error) {
	var resp *http.Response
	retryErr := retry.Do(
		func() error {
			// Get the access token
			storedToken, err := a.getAccessToken()
			if err != nil {
				return err
			}

			// Build the request using the access token
			req, err := requestBuilder(*storedToken)
			if err != nil {
				return err
			}

			// Attempt the request
			resp, err = a.httpClient.DoRequest(req)

			// If we run into an authentication error, refresh the access token
			if err == constants.ErrUnauthorized {
				tokenErr := a.executeAuthorizationWorkflow()
				if tokenErr != nil {
					return err
				}

				return err
			}

			return nil
		},
		retry.Attempts(3),
		retry.Delay(delay),
	)

	return resp, retryErr
}

func (a *API) getAccessToken() (*model.AccessToken, error) {
	storedToken, err := a.tokenStorage.GetFromFile()
	if err != nil {
		return nil, err
	}

	if storedToken == nil {
		err := a.executeAuthorizationWorkflow()
		if err != nil {
			return nil, err
		}

		return a.getAccessToken()
	}

	return storedToken, nil
}

func (a *API) executeAuthorizationWorkflow() error {
	if a.config == nil {
		return constants.ErrMissingConfig
	}

	authorizationCode, err := a.requestAuthorizationCode()
	if err != nil {
		return err
	}

	token, err := a.requestAccessToken(*authorizationCode)
	if err != nil {
		return err
	}

	return a.tokenStorage.WriteToFile(*token)
}

func (a *API) requestAuthorizationCode() (*string, error) {
	authorizeURL := convert.BuildAuthorizeURL(a.config.Authentication.Scope, a.config.App.ClientID, a.config.App.RedirectURI, a.config.App.State)

	err := browser.OpenURL(authorizeURL)
	if err != nil {
		return nil, err
	}

	var callbackURLString string
	fmt.Println("Enter callback URL: ")
	fmt.Scanln(&callbackURLString)

	authorizationCode, err := convert.AuthorizationCodeFromCallbackURL(callbackURLString, a.config.App.State)
	if err != nil {
		return nil, err
	}

	return authorizationCode, nil
}

func (a *API) requestAccessToken(authorizationCode string) (*model.AccessToken, error) {
	req, err := convert.BuildAccessTokenRequest(a.config.Authentication.GrantType, authorizationCode, a.config.App.RedirectURI, a.config.Authentication.ContentType, a.config.Authentication.Authorization, a.config.App.ClientID, a.config.App.ClientSecret)
	if err != nil {
		return nil, err
	}

	resp, err := a.executeAccessTokenRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	accessTokenResponse, err := convert.DecodeAccessTokenResponseBody(*resp)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponse.AccessToken, nil
}

func (a *API) executeAccessTokenRequest(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	retryErr := retry.Do(
		func() error {
			var err error

			// Attempt the request
			resp, err = a.httpClient.DoRequest(req)
			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(3),
		retry.Delay(delay),
	)

	return resp, retryErr
}
