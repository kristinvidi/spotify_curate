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

	ID     int32 `bun:"id,pk,autoincrement"`
	UserID ID    `bun:"user_spotify_id,notnull"`
	Genre  Genre `bun:",notnull"`
}

type UserArtistIDGenreMapping struct {
	bun.BaseModel `bun:"table:user_artist_spotify_id_genre_mapping"`

	UserID   ID    `bun:"user_spotify_id,notnull"`
	ArtistID ID    `bun:"artist_spotify_id,notnull"`
	GenreID  int32 `bun:",notnull"`
}

type UserArtistIDGenreMappings []UserArtistIDGenreMapping

type ArtistAlbumIDMapping struct {
	bun.BaseModel `bun:"table:artist_album_spotify_id_mapping"`

	ArtistID ID `bun:"artist_spotify_id,notnull"`
	AlbumID  ID `bun:"album_spotify_id,notnull"`
}

type ArtistAlbumIDMappings []ArtistAlbumIDMapping

type UserPlaylistTrackIDMapping struct {
	bun.BaseModel `bun:"table:user_playlist_track_spotify_id_mapping"`

	UserID     ID `bun:"user_spotify_id,notnull"`
	PlaylistID ID `bun:"playlist_spotify_id,notnull"`
	TrackID    ID `bun:"track_spotify_id,notnull"`
}
