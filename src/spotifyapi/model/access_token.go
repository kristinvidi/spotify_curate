package model

import "fmt"

type AccessToken string

func (a *AccessToken) HeaderValue() string {
	return fmt.Sprintf("Bearer %s", *a)
}

type AccessTokenResponse struct {
	AccessToken  AccessToken `json:"access_token"`
	TokenType    string      `json:"token_type"`
	ExpiresIn    int         `json:"expires_in"`
	RefreshToken string      `json:"refresh_token"`
	Scope        string      `json:"scope"`
}
