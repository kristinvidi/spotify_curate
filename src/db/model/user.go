package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:spotify_user"`

	ID        ID        `bun:"spotify_id,unique,notnull"`
	URI       URI       `bun:"uri,unique,notnull"`
	Name      string    `bun:"display_name,notnull"`
	Email     string    `bun:",notnull"`
	Country   string    `bun:",notnull"`
	CreatedAt time.Time `bun:",notnull"`
}

type UserUpdateStatus struct {
	bun.BaseModel `bun:"table:user_update_status"`

	UserID    ID        `bun:"user_spotify_id,notnull"`
	UpdatedAt time.Time `bun:",notnull"`
}

type UserSavedTracks struct {
	bun.BaseModel `bun:"table:user_saved_tracks"`

	UserID  ID        `bun:"user_spotify_id,notnull"`
	TrackID ID        `bun:"track_spotify_id,notnull"`
	SavedAt time.Time `bun:",notnull"`
}
