package domain

import (
	"strings"

	"src/config"
	"src/db/query"
	"src/domain/mapper"
	"src/domain/model"
	"src/spotifyapi"
	"src/spotifyapi/api"

	"go.openly.dev/pointy"
	"go.uber.org/zap"
)

type UserManager struct {
	config  *config.Config
	userAPI *api.User
	db      *query.PostgresDB
	logger  *zap.Logger
}

func NewUserManager(config *config.Config, logger *zap.Logger) *UserManager {
	return &UserManager{
		config:  config,
		userAPI: spotifyapi.GetUser(config),
		db:      query.NewPostgresDB(config.Database),
		logger:  logger,
	}
}

func (u *UserManager) AuthenticateUser() (*string, error) {
	user, err := u.getAndStoreCurrentUserProfile()
	if err != nil {
		return nil, err
	}

	return pointy.String(string(user.ID)), nil
}

func (u *UserManager) UpdateUserData() error {
	user, err := u.getAndStoreCurrentUserProfile()
	if err != nil {
		return err
	}

	tracksResponses, err := u.userAPI.GetUsersSavedTracks()
	if err != nil {
		return err
	}

	err = u.db.UpsertUserSavedTracks(
		mapper.IDToDBID(user.ID),
		mapper.DBUserSavedTracksFromGetUsersSavedTracksResponse(tracksResponses, user.ID),
	)
	if err != nil {
		return err
	}

	artistsResponses, err := u.userAPI.GetCurrentUsersFollowedArtists()
	if err != nil {
		return err
	}

	// Insert artist data
	dbArtists := mapper.DBFollowedArtistsFromGetFollowedArtistsResponse(artistsResponses)

	err = u.db.InsertArtistData(dbArtists)
	if err != nil {
		return err
	}

	// Remove unfollowed artists genre mappings
	rowsDeleted, err := u.db.DeleteUserToArtistIDGenreMappings(mapper.IDToDBID(user.ID), dbArtists.IDs())
	if err != nil {
		return err
	}

	u.logger.Info("successfully removed user artist genre mappings", zap.Int64("rowsDeleted", rowsDeleted))

	// Remove unfollowed artists mappings
	rowsDeleted, err = u.db.DeleteUserToArtistIDMappings(mapper.IDToDBID(user.ID), dbArtists.IDs())
	if err != nil {
		return err
	}

	u.logger.Info("successfully removed user artist mappings", zap.Int64("rowsDeleted", rowsDeleted))

	// Insert user to artist mapping data
	err = u.db.InsertUserToArtistIDMappings(
		mapper.DBUserToArtistMappingFromGetFollowedArtistsResponse(user.ID, artistsResponses),
	)
	if err != nil {
		return err
	}

	artistIDToAlbumsResponses, err := u.userAPI.GetCurrentUsersFollowedArtistsToAlbums(
		mapper.APIArtistsFromGetFollowedArtistsResponse(artistsResponses),
	)
	if err != nil {
		return err
	}

	albums, artistToAlbumMappings := mapper.DBAAbumsAndArtistMappingFromGetArtistsAlbumsResponse(artistIDToAlbumsResponses)

	err = u.db.InsertAlbums(albums)
	if err != nil {
		return err
	}

	err = u.db.InsertArtistAlbumIDMappings(artistToAlbumMappings)
	if err != nil {
		return err
	}

	err = u.db.InsertUserUpdateStatus(
		mapper.UserUpdateStatus(user.ID),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserManager) GetUnmappedArtistsForUser(spotifyUserID string) ([]model.Artist, error) {
	dbUserID := mapper.StringToDBID(spotifyUserID)

	artists, err := u.db.GetUnmappedArtistsForUser(dbUserID)
	if err != nil {
		return nil, err
	}

	return mapper.ArtistsFromDBArtists(artists), nil
}

func (u *UserManager) CreateArtistToGenreMappingForUser(spotifyUserID string, mappings []model.GenreToArtistsMapping) ([]model.GenreToArtistsMapping, error) {
	var unfollowedArtists []model.GenreToArtistsMapping

	dbUserID := mapper.StringToDBID(spotifyUserID)
	for _, mapping := range mappings {
		genre, err := u.db.GetGenreMappingForUserAndGenreName(dbUserID, mapping.Genre)
		if err != nil {
			return nil, err
		}

		dbArtists, err := u.db.GetMappedArtistsForUserByArtistNames(dbUserID, mapping.ArtistNames)
		if err != nil {
			return nil, err
		}

		artists := mapper.ArtistsFromDBArtists(dbArtists)
		unfollowed := u.getUnfollowedArtistNames(artists, mapping.ArtistNames)

		if len(unfollowed) > 0 {
			genreToArtistMapping := model.GenreToArtistsMapping{Genre: mapping.Genre, ArtistNames: unfollowed}

			unfollowedArtists = append(unfollowedArtists, genreToArtistMapping)
		}

		genreToArtistMappings := mapper.DBUserArtistIDGenreMappingsFromDBGenreAndArtists(*genre, artists)

		err = u.db.InsertUserArtistIDGenreMappings(genreToArtistMappings)
		if err != nil {
			return nil, err
		}
	}

	return unfollowedArtists, nil
}

func (u *UserManager) getUnfollowedArtistNames(artists []model.Artist, artistNames []string) []string {
	var unfollowed []string
	for _, artistName := range artistNames {
		if !u.artistNameInArtists(artistName, artists) {
			unfollowed = append(unfollowed, artistName)
		}
	}

	return unfollowed
}

func (u *UserManager) artistNameInArtists(artistName string, artists []model.Artist) bool {
	for _, artist := range artists {
		if artist.Name == artistName {
			return true
		}
	}

	return false
}

func (u *UserManager) getAndStoreCurrentUserProfile() (*model.User, error) {
	response, err := u.userAPI.GetCurrentUsersProfile()
	if err != nil {
		return nil, err
	}

	user := mapper.UserFromCurrentUsersProfileResponse(response)

	err = u.db.InsertUserData(
		*mapper.UserToDBUser(user),
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserManager) CreateLabelsForUser(spotifyUserID string, labels []string) ([]string, error) {
	var failedLabels []string
	dbUserID := mapper.StringToDBID(spotifyUserID)

	mapping := mapper.DBUserIDGenreMappingFromUserIDAndLabel(dbUserID, labels)

	err := u.db.InsertUserIDGenreMappings(mapping)
	if err != nil {
		u.logger.Error("failed to create labels for user",
			zap.String("user_id", spotifyUserID),
			zap.String("labels", strings.Join(labels, ", ")),
			zap.Error(err))
	}

	return failedLabels, err
}
