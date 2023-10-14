package server

import (
	"fmt"
	"src/config"
	"src/domain"
	"time"

	"github.com/sirupsen/logrus"
)

type GrpcServer struct {
	config *config.Config
	logger logrus.FieldLogger
}

func NewGrpcServer(config *config.Config, logger logrus.FieldLogger) *GrpcServer {
	return &GrpcServer{
		config: config,
		logger: logger,
	}
}

func (g *GrpcServer) UpdateUserData() error {
	api := "update_user_data"
	fmt.Printf("calling api: %s\n", api)

	updater := domain.NewUserUpdater(g.config)

	err := updater.UpdateUserData()
	if err != nil {
		return err
	}

	fmt.Println("successfully updated user data")

	return nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenre(genre string, relativeDate time.Time) error {
	api := "create_recent_in_genre"
	fmt.Printf("calling api: %s\n", api)

	creator := domain.NewPlaylistCreator(g.config)

	err := creator.CreateRecentInGenre(genre, relativeDate)
	if err != nil {
		return err
	}

	fmt.Printf("successfully created %s playlist\n", genre)

	return nil
}
