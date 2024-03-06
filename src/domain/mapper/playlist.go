package mapper

import (
	"time"

	db "src/db/model"
)

func DBPlaylistRecentInGenreGeneratedStatus(userID, playlistID db.ID, genreID int32) db.PlaylistRecentInGenreGeneratedStatus {
	return db.PlaylistRecentInGenreGeneratedStatus{
		UserID:     userID,
		PlaylistID: playlistID,
		GenreID:    genreID,
		CreatedAt:  time.Now(),
	}
}
