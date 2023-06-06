package httprequest

import (
	"fmt"

	"github.com/pkg/browser"

	"spotify_app/api/config"
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/converter"
)

func (h *HttpRequest) GetAccessToken(config config.Config) (*apptype.AccessToken, error) {
	existingToken, err := h.configManager.GetAccessToken()
	if err != nil {
		return nil, err
	}

	tokenValid, err := h.existingTokenIsValid(*existingToken)
	if err != nil {
		return nil, err
	}

	switch *tokenValid {
	case true:
		fmt.Println("Happy token yay!")
		return existingToken, nil
	default:
		fmt.Println("Getting new token")
		return h.getAccessToken(config)
	}
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

func (h *HttpRequest) existingTokenIsValid(existingToken apptype.AccessToken) (*bool, error) {
	request, err := converter.BuildGetCurrentUsersProfileRequest(existingToken)
	if err != nil {
		return nil, err
	}

	badRequest, err := h.checkResponseStatusIsBadRequest(request)
	if err != nil {
		return nil, err
	}

	tokenIsValid := !*badRequest

	return &tokenIsValid, nil
}

func (h *HttpRequest) requestAccessToken(config config.Config, authorizationCode string) (*apptype.AccessToken, error) {
	request, err := converter.BuildAccessTokenRequest(config.Authentication.GrantType, authorizationCode, config.App.RedirectURI, config.Authentication.ContentType, config.Authentication.Authorization, config.App.ClientID, config.App.ClientSecret)
	if err != nil {
		return nil, err
	}

	response, err := h.DoRequestAndCheckStatusIsOK(request)
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

func (h *HttpRequest) getAccessToken(config config.Config) (*apptype.AccessToken, error) {
	authorizationCode, err := h.requestAuthorizationCode(config)
	if err != nil {
		return nil, err
	}

	token, err := h.requestAccessToken(config, *authorizationCode)
	if err != nil {
		return nil, err
	}

	err = h.configManager.WriteAccessTokenToFile(*token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
