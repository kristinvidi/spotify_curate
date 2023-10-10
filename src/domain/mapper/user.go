package mapper

import (
	db "src/db/model"
	"src/domain/model"
	api "src/spotifyapi/model"
	"time"
)

func UserFromCurrentUsersProfileResponse(response *api.GetCurrentUsersProfileResponse) *model.User {
	if response == nil {
		return nil
	}

	return &model.User{
		DisplayName: model.Name(response.DisplayName),
		Email:       model.Email(response.Email),
		ID:          model.ID(response.ID),
		URI:         model.URI(response.URI),
		Country:     model.CountryCode(response.Country),
	}
}

func UserToDBUser(user *model.User) *db.User {
	if user == nil {
		return nil
	}

	return &db.User{
		DisplayName: string(user.DisplayName),
		Email:       string(user.Email),
		ID:          string(user.ID),
		URI:         string(user.URI),
		Country:     string(user.Country),
		CreatedAt:   time.Now(),
	}
}

func DBUserToArtistMappingFromGetFollowedArtistsResponse(userID model.ID, response []*api.GetFollowedArtistsResponse) []db.UserArtistSpotifyIDMapping {
	var dbUserArtistSpotifyIDMappings []db.UserArtistSpotifyIDMapping

	for _, r := range response {
		if r == nil {
			continue
		}

		for _, a := range r.Artists.ArtistList {
			mapping := db.UserArtistSpotifyIDMapping{
				UserID:    string(userID),
				ArtistID:  string(a.ID),
				CreatedAt: time.Now(),
			}

			dbUserArtistSpotifyIDMappings = append(dbUserArtistSpotifyIDMappings, mapping)
		}
	}

	return dbUserArtistSpotifyIDMappings
}

func APIArtistsFromGetFollowedArtistsResponse(response []*api.GetFollowedArtistsResponse) api.Artists {
	var artists api.Artists

	for _, r := range response {
		if r == nil {
			continue
		}

		artists = append(artists, r.Artists.ArtistList...)
	}

	return artists
}

func DBFollowedArtistsFromGetFollowedArtistsResponse(response []*api.GetFollowedArtistsResponse) []db.Artist {
	var dbArtists []db.Artist

	for _, r := range response {
		if r == nil {
			continue
		}

		for _, a := range r.Artists.ArtistList {
			dbArtist := db.Artist{
				SpotifyID: string(a.ID),
				URI:       string(a.URI),
				Name:      string(a.Name),
				CreatedAt: time.Now(),
			}

			dbArtists = append(dbArtists, dbArtist)
		}
	}

	return dbArtists
}
