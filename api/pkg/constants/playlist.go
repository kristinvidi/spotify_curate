package constants

const (
	PlaylistName          string = "name"
	PlaylistDescription   string = "description"
	PlaylistPublic        string = "public"
	PlaylistCollaborative string = "collaborative"
)

type PlaylistType int32

const (
	TopTracksForArtistsPlaylist PlaylistType = iota
	RecentTracksForArtistsPlaylist
)
