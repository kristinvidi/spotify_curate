package constants

const (
	HeaderAuthorization string = "Authorization"
	HeaderContentType   string = "Content-Type"
)

const (
	URLScheme                  string = "https"
	URLHostAccounts            string = "accounts.spotify.com"
	URLHostAPI                 string = "api.spotify.com"
	URLPathAuthorize           string = "authorize"
	URLPathToken               string = "api/token"
	URLPathMe                  string = "v1/me"
	URLPathMeFollowing         string = "v1/me/following"
	URLPathTrack               string = "v1/tracks"
	URLPathArtist              string = "v1/artists"
	URLPathUsers               string = "v1/users"
	URLPathPlaylists           string = "v1/playlists"
	URLSpecifierTopTracks      string = "top-tracks"
	URLSpecifierRelatedArtists string = "related-artists"
	URLSpecifierAlbums         string = "albums"
	URLSpecifierPlaylists      string = "playlists"
	URLSpecifierTracks         string = "tracks"
)

const (
	ParameterType   string = "type"
	ParameterLimit  string = "limit"
	ParameterAfter  string = "after"
	ParameterMarket string = "market"
)

const (
	TypeArtist string = "artist"
)

const (
	MarketES string = "ES"
)

const (
	ReplyLimit int = 50
)
