package spotifyapi

import (
	"src/config"
	"src/spotifyapi/api"
	"src/spotifyapi/authentication"
	"src/spotifyapi/client"
	"src/spotifyapi/convert"
)

func GetUser(c *config.Config) *api.User {
	return api.NewUser(
		api.NewAPI(client.NewHttp(), c, authentication.NewAccessTokenStorage()),
		convert.NewUser(),
	)
}
