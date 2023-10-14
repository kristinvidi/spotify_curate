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
		convert.NewArtist(),
	)
}

func GetArtist(c *config.Config) *api.Artist {
	return api.NewArtist(
		api.NewAPI(client.NewHttp(), c, authentication.NewAccessTokenStorage()),
		convert.NewArtist(),
	)
}

func GetAlbum(c *config.Config) *api.Album {
	return api.NewAlbum(
		api.NewAPI(client.NewHttp(), c, authentication.NewAccessTokenStorage()),
		convert.NewAlbum(),
	)
}

func GetPlaylist(c *config.Config) *api.Playlist {
	return api.NewPlaylist(
		api.NewAPI(client.NewHttp(), c, authentication.NewAccessTokenStorage()),
		convert.NewPlaylist(),
	)
}
