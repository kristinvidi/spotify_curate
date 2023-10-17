package server

import (
	"src/config"
	"src/domain"

	"go.uber.org/zap"
)

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

func (g *GrpcServer) UpdateUserData() error {
	api := "update_user_data"
	g.logger.Info("calling api", zap.String("api", api))

	updater := domain.NewUserUpdater(g.config)

	err := updater.UpdateUserData()
	if err != nil {
		return err
	}

	g.logger.Info("successfully updated user data")

	return nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenre(genre string) error {
	api := "create_playlist_recent_in_genre"
	g.logger.Info("calling api", zap.String("api", api))

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
	g.logger.Info("calling api", zap.String("api", api))

	creator := domain.NewPlaylistCreator(g.config)

	count, err := creator.CreateRecentInGenreAll()
	if err != nil {
		g.logger.Error("failure occured while creating recent in genre playlists for mapped genres", zap.Error(err), zap.Int("playlistsCreated", count))
		return err
	}

	g.logger.Info("successfully created recent in genre playlists for mapped genres", zap.Int("playlistCount", count))

	return nil
}
