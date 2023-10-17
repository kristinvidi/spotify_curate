package main

import (
	"fmt"
	"src/config"
	"src/server"

	"go.uber.org/zap"
)

type Job int32

const (
	UPDATE_USER_DATA Job = iota
	CREATE_PLAYLIST_RECENT_IN_GENRE
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := zap.Must(zap.NewDevelopment())
	if config.AppEnv.Env == "production" {
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync()

	logger.Info("Starting Curate!")

	grpcServer := server.NewGrpcServer(config, logger)

	job := CREATE_PLAYLIST_RECENT_IN_GENRE

	switch job {
	case UPDATE_USER_DATA:
		err = grpcServer.UpdateUserData()

	case CREATE_PLAYLIST_RECENT_IN_GENRE:
		genre := "Minimal Techno"
		err = grpcServer.CreatePlaylistRecentInGenre(genre)
	}

	if err != nil {
		fmt.Println(err)
	}
}
