package converter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"spotify_app/api/pkg/constants"
	"spotify_app/api/pkg/model"
)

func BuildAuthorizeURL(authenticationScope, clientID, redirectURI, appState string) string {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAccounts,
		Path:   constants.URLPathAuthorize,
	}
	params := url.Query()
	params.Set(constants.AuthenticationResponseType, constants.AuthenticationCode)
	params.Set(constants.AuthenticationScope, authenticationScope)
	params.Set(constants.AuthenticationClientID, clientID)
	params.Set(constants.AuthenticationRedirectURI, redirectURI)
	params.Set(constants.AuthenticationState, appState)

	url.RawQuery = params.Encode()

	return url.String()
}

func BuildAccessTokenRequest(grantType, authCode, redirectURI, contentTypeHeader, authorizationHeader, clientID, clientSecret string) (*http.Request, error) {
	url := url.URL{
		Scheme: constants.URLScheme,
		Host:   constants.URLHostAccounts,
		Path:   constants.URLPathToken,
	}
	baseURL := url.String()

	params := url.Query()
	params.Set(constants.AuthenticationGrantType, grantType)
	params.Set(constants.AuthenticationCode, authCode)
	params.Set(constants.AuthenticationRedirectURI, redirectURI)
	query := params.Encode()

	req, err := http.NewRequest(http.MethodPost, baseURL, strings.NewReader(query))
	if err != nil {
		return nil, err
	}

	req.Header.Add(constants.HeaderContentType, contentTypeHeader)
	req.Header.Add(constants.HeaderAuthorization, encodeAuthorization(authorizationHeader, clientID, clientSecret))

	return req, nil
}

func AuthenticationCodeFromCallbackURL(callbackURLString, configState string) (*string, error) {
	callbackURL, err := url.Parse(callbackURLString)
	if err != nil {
		return nil, err
	}

	params, err := url.ParseQuery(callbackURL.RawQuery)
	if err != nil {
		return nil, err
	}

	err = stateIsInURLParams(params, configState)
	if err != nil {
		return nil, err
	}

	code, err := codeFromURLParams(params)
	if err != nil {
		return nil, err
	}

	return code, nil
}

func stateIsInURLParams(urlParams url.Values, configState string) error {
	state, ok := urlParams[constants.AuthenticationState]
	if !ok {
		return fmt.Errorf("state not found in callback URL")
	}

	if state[0] != configState {
		return fmt.Errorf("state value does not match config")
	}

	return nil
}

func codeFromURLParams(urlParams url.Values) (*string, error) {
	code, ok := urlParams[constants.AuthenticationCode]
	if !ok {
		return nil, fmt.Errorf("code not found in callback URL")
	}

	return &code[0], nil
}

func encodeAuthorization(authorizationType, clientID, clientSecret string) string {
	unencoded := fmt.Sprintf("%s:%s", clientID, clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(unencoded))

	return fmt.Sprintf("%s %s", authorizationType, encoded)
}

func DecodeAccessTokenResponseBody(res http.Response) (*model.AccessTokenResponse, error) {
	var accessTokenResponse model.AccessTokenResponse
	err := json.NewDecoder(res.Body).Decode(&accessTokenResponse)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponse, nil
}
