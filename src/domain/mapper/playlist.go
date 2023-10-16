package mapper

import (
	db "src/db/model"
	"time"
)

func DBPlaylistRecentInGenreGeneratedStatus(userID db.ID, genreID int32) db.PlaylistRecentInGenreGeneratedStatus {
	return db.PlaylistRecentInGenreGeneratedStatus{
		UserID:    userID,
		GenreID:   genreID,
		CreatedAt: time.Now(),
	}
}
