package constants

type Column string

const (
	ColumnID                  Column = "spotify_id"
	ColumnUserID              Column = "user_spotify_id"
	ColumnArtistID            Column = "artist_spotify_id"
	ColumnAlbumID             Column = "album_spotify_id"
	ColumnUserArtistID        Column = "user_spotify_id, artist_spotify_id"
	ColumnUserArtistGenreID   Column = "user_spotify_id, artist_spotify_id, genre_id"
	ColumnArtistAlbumID       Column = "artist_spotify_id, album_spotify_id"
	ColumnUserPlaylistTrackID Column = "user_spotify_id, playlist_spotify_id, track_spotify_id"
	ColumnUserTrackID         Column = "user_spotify_id, track_spotify_id"
	ColumnDisplayName         Column = "display_name"
	ColumnGenre               Column = "genre"
)

func (c Column) String() string {
	return string(c)
}
