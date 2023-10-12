package model

import (
	"time"

	"github.com/uptrace/bun"
)

type UserArtistIDMapping struct {
	bun.BaseModel `bun:"table:user_artist_spotify_id_mapping"`

	UserID    ID        `bun:"user_spotify_id,notnull"`
	ArtistID  ID        `bun:"artist_spotify_id,notnull"`
	CreatedAt time.Time `bun:",notnull"`
}

type UserIDGenreMapping struct {
	bun.BaseModel `bun:"table:user_spotify_id_genre_mapping"`

	UserID ID    `bun:"user_spotify_id,notnull"`
	Genre  Genre `bun:",notnull"`
}

type UserArtistIDGenreMapping struct {
	bun.BaseModel `bun:"table:user_artist_spotify_id_genre_mapping"`

	UserID   ID    `bun:"user_spotify_id,notnull"`
	ArtistID URI   `bun:"artist_spotify_id,notnull"`
	GenreID  int32 `bun:",notnull"`
}
