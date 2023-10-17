package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Album struct {
	bun.BaseModel `bun:"table:spotify_album"`

	ID                   ID        `bun:"spotify_id,unique,notnull"`
	URI                  URI       `bun:"uri,unique,notnull"`
	Name                 string    `bun:"display_name,notnull"`
	Type                 AlbumType `bun:"album_type,notnull"`
	ReleaseDate          time.Time
	ReleaseDatePrecision ReleaseDatePrecision
	CreatedAt            time.Time `bun:",notnull"`
}

type Albums []Album

type AlbumType string

const (
	AlbumTypeAlbum       AlbumType = "album"
	AlbumTypeSingle      AlbumType = "single"
	AlbumTypeCompilation AlbumType = "compilation"
)
