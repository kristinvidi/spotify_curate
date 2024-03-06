package mapper

import (
	db "src/db/model"
	"src/domain/model"
)

func DBUserArtistIDGenreMappingsFromDBGenreAndArtists(genre db.UserIDGenreMapping, artists model.Artists) db.UserArtistIDGenreMappings {
	mappings := make(db.UserArtistIDGenreMappings, len(artists))
	for i, artist := range artists {
		mappings[i] = db.UserArtistIDGenreMapping{
			UserID:   genre.UserID,
			ArtistID: IDToDBID(artist.ID),
			GenreID:  genre.ID,
		}
	}

	return mappings
}

func DBUserPlaylistTrackIDMappings(userID, playlistID db.ID, trackIDs []db.ID) []db.UserPlaylistTrackIDMapping {
	mappings := make([]db.UserPlaylistTrackIDMapping, len(trackIDs))

	for i, trackID := range trackIDs {
		mappings[i] = db.UserPlaylistTrackIDMapping{
			UserID:     userID,
			PlaylistID: playlistID,
			TrackID:    trackID,
		}
	}

	return mappings
}
