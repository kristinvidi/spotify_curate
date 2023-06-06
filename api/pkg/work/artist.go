package work

import (
	"fmt"

	apptype "spotify_app/api/pkg/app_type"
	datastore "spotify_app/api/pkg/data_store"
	httprequest "spotify_app/api/pkg/http_request"
	"spotify_app/api/pkg/model"
)

func StoreArtistInfo(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken) (model.Artists, error) {
	followedArtists, err := httpRequest.GetFollowedArtists(accessToken)
	if err != nil {
		return nil, err
	}

	artists, artistsWithoutGenres := sortArtistsWithAndWithoutGenres(followedArtists)

	if len(artistsWithoutGenres) == 0 {
		return artists, nil
	}

	for _, artist := range artistsWithoutGenres {
		fmt.Printf("getting genre for %s\n", artist.Name)
		genre, err := getGenreForArtist(httpRequest, accessToken, artist)
		if err != nil {
			return nil, err
		}
		if genre != nil {
			artist.SetGenres([]apptype.Genre{*genre})
		}

		artists = append(artists, artist)
	}

	db := datastore.NewArtistTextDB()

	err = db.WriteAllEntries(artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func sortArtistsWithAndWithoutGenres(artists model.Artists) (model.Artists, model.Artists) {
	artistsWithGenres := []model.Artist{}
	artistsWithoutGenres := []model.Artist{}

	for _, artist := range artists {
		if len(artist.Genres) == 0 {
			artistsWithoutGenres = append(artistsWithoutGenres, artist)
		} else {
			artistsWithGenres = append(artistsWithGenres, artist)
		}
	}

	return artistsWithGenres, artistsWithoutGenres
}

func getGenreForArtist(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken, artist model.Artist) (*apptype.Genre, error) {
	genreFromRelatedArtists, err := getGenreFromArtistsRelatedArtists(httpRequest, accessToken, artist)
	if err != nil {
		return nil, err
	}
	if genreFromRelatedArtists != nil {
		return genreFromRelatedArtists, nil
	}

	return getGenreFromArtistsTopTracks(httpRequest, accessToken, artist)
}

func getGenreFromArtistsRelatedArtists(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken, artist model.Artist) (*apptype.Genre, error) {
	relatedArtists, err := httpRequest.GetArtistRelatedArtists(accessToken, artist.ID)
	if err != nil {
		return nil, err
	}

	return relatedArtists.TopGenre(), nil
}

func getGenreFromArtistsTopTracks(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken, artist model.Artist) (*apptype.Genre, error) {
	topTracks, err := httpRequest.GetArtistTopTracks(accessToken, artist.ID)
	if err != nil {
		return nil, err
	}

	for _, track := range *topTracks {
		genre := track.Artists.TopGenre()
		if genre == nil {
			continue
		}

		return genre, nil
	}

	return nil, nil
}
