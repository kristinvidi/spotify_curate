package model

import (
	"time"

	"github.com/uptrace/bun"
)

type PlaylistRecentInGenreGeneratedStatus struct {
	bun.BaseModel `bun:"table:playlist_recent_in_genre_generated_status"`

	UserID     ID        `bun:"user_spotify_id,notnull"`
	PlaylistID ID        `bun:"playlist_spotify_id"`
	GenreID    int32     `bun:",notnull"`
	CreatedAt  time.Time `bun:",notnull"`
}
