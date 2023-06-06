package main

import (
	"fmt"
	"net/http"

	"spotify_app/api/config"
	"spotify_app/api/pkg/constants"
	httprequest "spotify_app/api/pkg/http_request"
	"spotify_app/api/pkg/work"
)

type Workflow int32

const (
	GetAccessToken Workflow = iota
	StoreArtistInfo
	StoreAlbumInfo
	CreatePlaylistForGenre
)

func main() {
	viperConfig, err := config.GetEnvironmentVariables()
	if err != nil {
		panic(err)
	}

	configManager := config.NewManager()
	httpRequest := httprequest.NewHttpRequest(&http.Client{}, configManager)

	accessToken, err := httpRequest.GetAccessToken(viperConfig)
	if err != nil {
		panic(err)
	}

	err = work.StoreUserInfo(httpRequest, *accessToken, configManager)
	if err != nil {
		panic(err)
	}

	switch CreatePlaylistForGenre {
	case GetAccessToken:
		fmt.Println("Successfully saved access token")
	case StoreArtistInfo:
		_, err = work.StoreArtistInfo(httpRequest, *accessToken)
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully stored artist info")
	case CreatePlaylistForGenre:
		genre := constants.GenreProgressiveTechHouse
		playlistType := constants.TopTracksForArtistsPlaylist

		err = work.CreatePlaylistForGenre(httpRequest, *accessToken, genre, playlistType)
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully created playlist")
	// case StoreAlbumInfo:

	default:
		fmt.Println("Did nothing! Method not yet implemented.")
	}
}
