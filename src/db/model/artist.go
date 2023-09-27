package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Artist struct {
	bun.BaseModel `bun:"table:artist"`

	ID        *int64    `bun:"id"`
	SpotifyID string    `bun:"spotify_id"`
	URI       string    `bun:"uri"`
	Name      string    `bun:"name"`
	CreatedAt time.Time `bun:"created_at"`
}
