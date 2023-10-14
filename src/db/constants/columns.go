package constants

type Column string

const (
	ColumnID            Column = "spotify_id"
	ColumnUserArtistID  Column = "user_spotify_id, artist_spotify_id"
	ColumnArtistAlbumID Column = "artist_spotify_id, album_spotify_id"
)
