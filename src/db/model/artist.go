package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Artist struct {
	bun.BaseModel `bun:"table:spotify_artist"`

	ID        ID        `bun:"spotify_id,unique,notnull"`
	URI       URI       `bun:"uri,unique,notnull"`
	Name      string    `bun:"display_name,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
}

type Artists []Artist

func (a *Artists) IDs() []ID {
	var ids []ID
	for _, artist := range *a {
		ids = append(ids, artist.ID)
	}

	return ids
}
