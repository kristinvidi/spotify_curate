package domain

import (
	"src/config"
	"src/db/query"
	"src/domain/mapper"
	"src/domain/model"
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

func (p *PlaylistCreator) CreateRecentInGenreAll() (int, error) {
	response, err := p.userAPI.GetCurrentUsersProfile()
	if err != nil {
		return 0, err
	}

	user := mapper.UserFromCurrentUsersProfileResponse(response)
	dbUser := mapper.UserToDBUser(user)

	// Get Genres for user
	mappings, err := p.db.GetGenreMappingsForUser(dbUser.ID)
	if err != nil {
		return 0, err
	}

	playlistCount := 0
	for _, m := range mappings {
		err = p.createRecentInGenrePlaylist(user, string(m.Genre))
		if err != nil {
			return playlistCount, err
		}

		playlistCount += 1
	}

	return playlistCount, nil
}

func (p *PlaylistCreator) CreateRecentInGenre(genre string) (bool, error) {
	response, err := p.userAPI.GetCurrentUsersProfile()
	if err != nil {
		return false, err
	}

	user := mapper.UserFromCurrentUsersProfileResponse(response)

	err = p.createRecentInGenrePlaylist(user, genre)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *PlaylistCreator) createRecentInGenrePlaylist(user *model.User, genre string) error {
	dbUser := mapper.UserToDBUser(user)

	// See if Genre is mapped
	genreMapping, err := p.db.GetGenreMappingForUserAndGenre(dbUser.ID, mapper.StringToDBGenre(genre))
	if err != nil {
		return err
	}

	// Fetch date that a playlist was last created, if it exists!
	lastCreatedDate, err := p.getRelativeDateForNewPlaylistInGenre(*user, genre)
	if err != nil {
		return err
	}

	// Get albumIDs created after relative date
	albumIDs, err := p.db.GetAlbumIDsForGenreAfterDate(dbUser.ID, mapper.StringToDBGenre(genre), *lastCreatedDate)
	if err != nil {
		return err
	}

	if len(albumIDs) == 0 {
		// log a message saying nothing was generated
		return nil
	}

	// Get tracks for albums
	trackResponses, err := p.albumAPI.GetAlbumTracks(mapper.DBIDsToAPIIDs(albumIDs))
	if err != nil {
		return err
	}

	trackURIs := mapper.TrackAPIURIsFromGetAlbumTracksResponses(trackResponses)

	// Create playlist
	playlistName := p.playlistNameForRecentInGenre(genre, *lastCreatedDate)
	playlistResponse, err := p.playlistAPI.CreatePlaylist(*mapper.DBUserToAPIUserID(dbUser), playlistName, false, false, "")
	if err != nil {
		return err
	}

	// Add tracks to playlist
	_, err = p.playlistAPI.AddTracksToPlaylist(playlistResponse.ID, trackURIs)
	if err != nil {
		return err
	}

	// Update db with newly created playlist
	err = p.db.InsertPlaylistRecentInGenreGeneratedStatus(
		mapper.DBPlaylistRecentInGenreGeneratedStatus(dbUser.ID, genreMapping.ID),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistCreator) getRelativeDateForNewPlaylistInGenre(user model.User, genre string) (*time.Time, error) {
	// To not overwhelm the user, we'll only go back as far as 2 months max
	earliestDate := time.Now().AddDate(0, -2, 0)

	// Fetch the db last created date
	lastCreatedDate, err := p.db.GetLastCreatedAtDateForPlaylistOfGenre(mapper.IDToDBID(user.ID), mapper.StringToDBGenre(genre))
	if err != nil {
		return nil, err
	}

	// If the playlist genre does not have a last created date or if it's prior to earlistDate, use earliestDate
	if lastCreatedDate == nil || lastCreatedDate.Before(earliestDate) {
		return &earliestDate, nil
	}

	// Otherwise return the last created date!
	return lastCreatedDate, nil
}

func (p *PlaylistCreator) playlistNameForRecentInGenre(genre string, lastCreatedDate time.Time) string {
	return "Curate: Recent " + genre + " (from " + lastCreatedDate.Format("2006-01-02") + ")"
}
