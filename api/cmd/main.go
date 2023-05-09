package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"spotify_app/api/config"
	"spotify_app/api/pkg/data"
	httprequest "spotify_app/api/pkg/http_request"
)

func main() {
	config, err := config.GetEnvironmentVariables()
	if err != nil {
		panic(err)
	}

	httpRequest := httprequest.NewHttpRequest(&http.Client{})

	accessToken, err := httpRequest.GetAccessToken(config)
	if err != nil {
		panic(err)
	}

	followedArtists, err := httpRequest.GetFollowedArtists(*accessToken)
	if err != nil {
		panic(err)
	}

	mapOfGenreToArtists := data.GetMapOfGenreToArtists(followedArtists)
	for genre, artists := range mapOfGenreToArtists {
		fmt.Printf("Genre: %s; Artists: %s\n", genre, commaSeparatedStrings(artists))
	}
}

func commaSeparatedStrings(strings []string) string {
	commaSeparated, _ := json.Marshal(strings)
	return string(commaSeparated)
}

// Make an empty access_token.txt file
// Use httpRequest.GetAccessToken to populate
// Make request with access token
// 		If it fails on a specific code (auth token no longer valid), fetch a new one
//		Otherwise it just works
