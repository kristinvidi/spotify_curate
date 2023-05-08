package httprequest

import (
	"fmt"

	"github.com/pkg/browser"

	"spotify_app/api/config"
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/converter"
)

func (h *HttpRequest) GetAccessToken(config config.Config) (*apptype.AccessToken, error) {
	authorizationCode, err := h.requestAuthorizationCode(config)
	if err != nil {
		return nil, err
	}

	accessToken, err := h.requestAccessToken(config, *authorizationCode)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (h *HttpRequest) requestAuthorizationCode(config config.Config) (*string, error) {
	authorizeURL := converter.BuildAuthorizeURL(config.Authentication.Scope, config.App.ClientID, config.App.RedirectURI, config.App.State)

	err := browser.OpenURL(authorizeURL)
	if err != nil {
		return nil, err
	}

	var callbackURLString string
	fmt.Println("Enter callback URL: ")
	fmt.Scanln(&callbackURLString)

	authenticationCode, err := converter.AuthenticationCodeFromCallbackURL(callbackURLString, config.App.State)
	if err != nil {
		return nil, err
	}

	return authenticationCode, nil
}

func (h *HttpRequest) requestAccessToken(config config.Config, authorizationCode string) (*apptype.AccessToken, error) {
	request, err := converter.BuildAccessTokenRequest(config.Authentication.GrantType, authorizationCode, config.App.RedirectURI, config.Authentication.ContentType, config.Authentication.Authorization, config.App.ClientID, config.App.ClientSecret)
	if err != nil {
		return nil, err
	}

	response, err := h.DoRequestAndCheckStatus(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	accessTokenResponse, err := converter.DecodeAccessTokenResponseBody(*response)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponse.AccessToken, nil
}
