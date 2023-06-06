package work

import (
	"fmt"
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/constants"
	datastore "spotify_app/api/pkg/data_store"
	httprequest "spotify_app/api/pkg/http_request"
	"spotify_app/api/pkg/model"
	"strings"
	"time"
)

type PlaylistType int32

func CreatePlaylistForGenre(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken, genre constants.Genre, playlistType constants.PlaylistType) error {
	// Get artists for genre from data store
	db := datastore.NewMappingTextDB()
	genreToArtist, err := db.ReadGenreToArtistsMapping()
	if err != nil {
		return err
	}

	artists, ok := genreToArtist[genre]
	if !ok {
		return fmt.Errorf("no stored artists for genre %s", genre.String())
	}

	var tracks *model.Tracks
	// Get top tracks for each artist in genre
	switch playlistType {
	default:
		tracks, err = topTracksForArtists(httpRequest, accessToken, artists)
	}
	if err != nil {
		return err
	}

	// Create playlist with tracks!
	playlistID, err := httpRequest.CreatePlaylist(accessToken, playlistName(genre))
	if err != nil {
		return err
	}

	err = httpRequest.AddTracksToPlaylist(accessToken, *playlistID, *tracks)
	if err != nil {
		return err
	}

	return nil
}

func topTracksForArtists(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken, artistIDs []string) (*model.Tracks, error) {
	var tracks model.Tracks
	for _, id := range artistIDs {
		t, err := httpRequest.GetArtistTopTracks(accessToken, id)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, *t...)
	}

	return &tracks, nil
}

func newTracksForArtists(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken, artistIDs []string) (*model.Tracks, error) {
	// loop through each artist from the data store
	// get albums for artist
	// filter albums according to start date

	return nil, nil
}

func playlistName(genre constants.Genre) string {
	name := fmt.Sprintf("Kiki Curated %s %s", genre.String(), time.Now().Format("2006-01-02"))

	return strings.ToLower(name)
}
