package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Album struct {
	bun.BaseModel `bun:"table:album"`

	SpotifyID            string `bun:",unique,notnull"`
	URI                  string `bun:",unique,notnull"`
	Name                 string `bun:",notnull"`
	ReleaseDate          time.Time
	ReleaseDatePrecision ReleaseDatePrecision
	CreatedAt            time.Time `bun:",notnull"`
}

type Albums []Album
