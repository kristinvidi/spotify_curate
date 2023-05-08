package model

import (
	apptype "spotify_app/api/pkg/app_type"
)

type AccessTokenResponse struct {
	AccessToken  apptype.AccessToken `json:"access_token"`
	TokenType    string              `json:"token_type"`
	ExpiresIn    int                 `json:"expires_in"`
	RefreshToken string              `json:"refresh_token"`
	Scope        string              `json:"scope"`
}
