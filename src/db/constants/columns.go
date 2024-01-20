package constants

type Column string

const (
	ColumnID            Column = "spotify_id"
	ColumnUserID        Column = "user_spotify_id"
	ColumnArtistID      Column = "artist_spotify_id"
	ColumnAlbumID       Column = "album_spotify_id"
	ColumnUserArtistID  Column = "user_spotify_id, artist_spotify_id"
	ColumnArtistAlbumID Column = "artist_spotify_id, album_spotify_id"
)
