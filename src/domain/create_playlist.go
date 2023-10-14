package domain

import (
	"src/config"
	"src/db/query"
	"src/domain/mapper"
	"src/spotifyapi"
	"src/spotifyapi/api"
	"time"
)

type PlaylistCreator struct {
	config      *config.Config
	userAPI     *api.User
	albumAPI    *api.Album
	playlistAPI *api.Playlist
	db          *query.PostgresDB
}

func NewPlaylistCreator(config *config.Config) *PlaylistCreator {
	return &PlaylistCreator{
		config:      config,
		userAPI:     spotifyapi.GetUser(config),
		albumAPI:    spotifyapi.GetAlbum(config),
		playlistAPI: spotifyapi.GetPlaylist(config),
		db:          query.NewPostgresDB(config.Database),
	}
}

func (p *PlaylistCreator) CreateRecentInGenre(genre string, relativeDate time.Time) error {
	response, err := p.userAPI.GetCurrentUsersProfile()
	if err != nil {
		return err
	}

	user := mapper.UserToDBUser(mapper.UserFromCurrentUsersProfileResponse(response))

	// Get albumIDs created after relative date
	albumIDs, err := p.db.GetAlbumIDsForGenreAfterDate(user.ID, mapper.StringToDBGenre(genre), relativeDate)
	if err != nil {
		return err
	}

	// Get tracks for albums
	trackResponses, err := p.albumAPI.GetAlbumTracks(mapper.DBIDsToAPIIDs(albumIDs))
	if err != nil {
		return err
	}

	trackURIs := mapper.TrackAPIURIsFromGetAlbumTracksResponses(trackResponses)

	// Create playlist
	playlistName := "Curate: Recent " + genre + " (from " + relativeDate.Format("2006-01-02") + ")"
	playlistResponse, err := p.playlistAPI.CreatePlaylist(*mapper.DBUserToAPIUserID(user), playlistName, false, false, "")
	if err != nil {
		return err
	}

	// Add tracks to playlist
	_, err = p.playlistAPI.AddTracksToPlaylist(playlistResponse.ID, trackURIs)
	if err != nil {
		return err
	}

	return nil
}
