package api

import (
	"context"
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

func (a *API) DoRequest(requestBuilder func(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error), inputs *model.RequestInput) (*http.Response, error) {
	if a.config == nil {
		return nil, constants.ErrMissingConfig
	}

	resp, err := a.executeRequest(requestBuilder, inputs)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *API) executeRequest(requestBuilder func(accessToken model.AccessToken, inputs *model.RequestInput) (*http.Request, error), inputs *model.RequestInput) (*http.Response, error) {
	var resp *http.Response
	retryErr := retry.Do(
		func() error {
			// Get the access token
			storedToken, err := a.getAccessToken()
			if err != nil {
				return err
			}

			// Build the request using the access token
			req, err := requestBuilder(*storedToken, inputs)
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

			// If response is nil and error is nil, return missing response error
			if resp == nil && err == nil {
				return constants.ErrMissingResponse
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
	authorizeURL := convert.BuildAuthorizeURL(a.config.Authentication.Scope, a.config.AppClientInfo.ClientID, a.config.AppClientInfo.RedirectURI, a.config.AppClientInfo.State)

	// HTTP HANDLER HERE TO CAPTURE THE CALLBACK CODE
	var authorizationCode *string
	var serverErr error

	// Create a new server and set the handler
	server := http.Server{Addr: ":8888"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get authorization code from callback URL
		authorizationCode, serverErr = convert.AuthorizationCodeFromCallbackURL(r.URL.RawQuery, a.config.AppClientInfo.State)

		// Send a response to the user
		w.Write([]byte("Authorization complete. You can now close this window."))

		// Shut down the server
		server.Shutdown(context.Background())
	})

	// Start the server
	go server.ListenAndServe()

	// Open the browser to prompt user to login
	err := browser.OpenURL(authorizeURL)
	if err != nil {
		return nil, err
	}

	// We need a lil delay to let the program catch up with itself
	time.Sleep(6 * time.Second)

	// Wait for the server to shut down
	serverErr = server.ListenAndServe()
	if serverErr != nil && serverErr != http.ErrServerClosed {
		return nil, serverErr
	}

	return authorizationCode, nil
}

func (a *API) requestAccessToken(authorizationCode string) (*model.AccessToken, error) {
	req, err := convert.BuildAccessTokenRequest(a.config.Authentication.GrantType, authorizationCode, a.config.AppClientInfo.RedirectURI, a.config.Authentication.ContentType, a.config.Authentication.Authorization, a.config.AppClientInfo.ClientID, a.config.AppClientInfo.ClientSecret)
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
