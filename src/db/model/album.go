package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Album struct {
	bun.BaseModel `bun:"table:spotify_album"`

	ID                   ID     `bun:"spotify_id,unique,notnull"`
	URI                  URI    `bun:"uri,unique,notnull"`
	Name                 string `bun:"display_name,notnull"`
	ReleaseDate          time.Time
	ReleaseDatePrecision ReleaseDatePrecision
	CreatedAt            time.Time `bun:",notnull"`
}

type Albums []Album
