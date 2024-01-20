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
	GET_UNMAPPED_ARTISTS_FOR_USER
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

	job := GET_UNMAPPED_ARTISTS_FOR_USER

	switch job {
	case UPDATE_USER_DATA:
		err = grpcServer.UpdateUserData()

	case CREATE_PLAYLIST_RECENT_IN_GENRE:
		genre := "Tech House"
		err = grpcServer.CreatePlaylistRecentInGenre(genre)

	case CREATE_PLAYLIST_RECENT_IN_GENRE_ALL:
		err = grpcServer.CreatePlaylistRecentInGenreAll()

	case GET_UNMAPPED_ARTISTS_FOR_USER:
		err = grpcServer.GetUnmappedArtistsForUser()
	}

	if err != nil {
		fmt.Println(err)
	}
}
