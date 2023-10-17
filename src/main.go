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
	CREATE_PLAYLIST_RECENT_IN_GENRE_ALL
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

	job := CREATE_PLAYLIST_RECENT_IN_GENRE_ALL

	switch job {
	case UPDATE_USER_DATA:
		err = grpcServer.UpdateUserData()

	case CREATE_PLAYLIST_RECENT_IN_GENRE:
		genre := "Tech House"
		err = grpcServer.CreatePlaylistRecentInGenre(genre)

	case CREATE_PLAYLIST_RECENT_IN_GENRE_ALL:
		err = grpcServer.CreatePlaylistRecentInGenreAll()
	}

	if err != nil {
		fmt.Println(err)
	}
}
