package mapper

import (
	"time"

	db "src/db/model"
	"src/domain/model"
	api "src/spotifyapi/model"
)

func UserFromCurrentUsersProfileResponse(response *api.GetCurrentUsersProfileResponse) *model.User {
	if response == nil {
		return nil
	}

	return &model.User{
		Name:    string(response.DisplayName),
		Email:   string(response.Email),
		ID:      model.ID(response.ID),
		URI:     model.URI(response.URI),
		Country: string(response.Country),
	}
}

func UserToDBUser(user *model.User) *db.User {
	if user == nil {
		return nil
	}

	return &db.User{
		Name:      string(user.Name),
		Email:     string(user.Email),
		ID:        db.ID(user.ID),
		URI:       db.URI(user.URI),
		Country:   string(user.Country),
		CreatedAt: time.Now(),
	}
}

func DBUserToAPIUserID(user *db.User) *api.ID {
	if user == nil {
		return nil
	}

	id := api.ID(user.ID)
	return &id
}

func DBUserToArtistMappingFromGetFollowedArtistsResponse(userID model.ID, response []*api.GetFollowedArtistsResponse) []db.UserArtistIDMapping {
	var dbUserArtistIDMappings []db.UserArtistIDMapping

	for _, r := range response {
		if r == nil {
			continue
		}

		for _, a := range r.Artists.ArtistList {
			dbUserArtistIDMappings = append(dbUserArtistIDMappings, dbUserArtistIDMappingFromUserIDAndAPIArtist(userID, a))
		}
	}

	return dbUserArtistIDMappings
}

func dbUserArtistIDMappingFromUserIDAndAPIArtist(userID model.ID, artist api.Artist) db.UserArtistIDMapping {
	return db.UserArtistIDMapping{
		UserID:    db.ID(userID),
		ArtistID:  db.ID(artist.ID),
		CreatedAt: time.Now(),
	}
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

func DBFollowedArtistsFromGetFollowedArtistsResponse(response []*api.GetFollowedArtistsResponse) db.Artists {
	var dbArtists []db.Artist

	for _, r := range response {
		if r == nil {
			continue
		}

		for _, a := range r.Artists.ArtistList {
			dbArtists = append(dbArtists, dbArtistFromAPIArtist(a))
		}
	}

	return dbArtists
}

func dbArtistFromAPIArtist(artist api.Artist) db.Artist {
	return db.Artist{
		ID:        db.ID(artist.ID),
		URI:       db.URI(artist.URI),
		Name:      string(artist.Name),
		CreatedAt: time.Now(),
	}
}

func UserUpdateStatus(userID model.ID) db.UserUpdateStatus {
	return db.UserUpdateStatus{
		UserID:    db.ID(userID),
		UpdatedAt: time.Now(),
	}
}

func DBUserSavedTracksFromGetUsersSavedTracksResponse(response []*api.GetUsersSavedTracksResponse, userID model.ID) []db.UserSavedTracks {
	var userSavedTracks []db.UserSavedTracks

	for _, r := range response {
		if r == nil {
			continue
		}

		for _, a := range r.Items {
			track := db.UserSavedTracks{
				UserID:  db.ID(userID),
				TrackID: db.ID(a.Track.ID),
				SavedAt: a.SavedAt(),
			}
			userSavedTracks = append(userSavedTracks, track)
		}
	}

	return userSavedTracks
}
