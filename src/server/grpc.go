package server

import (
	"src/config"
	"src/domain"
	"src/domain/mapper"
	"src/server/converter"
	pb "src/server/proto"

	"go.uber.org/zap"
)

// TO-DO: Add grpc server logic, currently it's just an abstracted layer to facilitate building out the grpc code
type GrpcServer struct {
	config *config.Config
	logger *zap.Logger
}

func NewGrpcServer(config *config.Config, logger *zap.Logger) *GrpcServer {
	return &GrpcServer{
		config: config,
		logger: logger,
	}
}

func (g *GrpcServer) logAPICall(apiName string) {
	g.logger.Info("calling api", zap.String("api", apiName))
}

func (g *GrpcServer) UpdateUserData() (*pb.UpdateUserDataResponse, error) {
	api := "update_user_data"
	g.logAPICall(api)

	updater := domain.NewUserUpdater(g.config, g.logger)

	err := updater.UpdateUserData()
	if err != nil {
		return converter.SerializeUpdateUserData(false), err
	}

	g.logger.Info("successfully updated user data")

	return converter.SerializeUpdateUserData(true), nil
}

func (g *GrpcServer) GetUnmappedArtistsForUser() (*pb.GetUnmappedArtistsForUserResponse, error) {
	api := "get_unmapped_artists_for_user"
	g.logAPICall(api)

	updater := domain.NewUserUpdater(g.config, g.logger)

	artists, err := updater.GetUnmappedArtistsForUser()
	if err != nil {
		return converter.SerializeGetUnmappedArtistsForUser(false, nil), err
	}

	g.logger.Info("successfully fetched unmapped artists for user")

	return converter.SerializeGetUnmappedArtistsForUser(
		true,
		mapper.ServerArtistsFromDomainArtists(artists),
	), nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenre(genre string) error {
	api := "create_playlist_recent_in_genre"
	g.logAPICall(api)

	creator := domain.NewPlaylistCreator(g.config)

	generated, err := creator.CreateRecentInGenre(genre)
	if err != nil {
		return err
	}

	if !generated {
		g.logger.Info("no new content to add to playlist, skipping generating recent in genre playlist", zap.String("genre", genre))
		return nil
	}

	g.logger.Info("successfully created recent in genre playlist", zap.String("genre", genre))

	return nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenreAll() error {
	api := "create_playlist_recent_in_genre_all"
	g.logAPICall(api)

	creator := domain.NewPlaylistCreator(g.config)

	count, err := creator.CreateRecentInGenreAll()
	if err != nil {
		g.logger.Error("failure occured while creating recent in genre playlists for mapped genres", zap.Error(err), zap.Int("playlistsCreated", count))
		return err
	}

	g.logger.Info("successfully created recent in genre playlists for mapped genres", zap.Int("playlistCount", count))

	return nil
}
