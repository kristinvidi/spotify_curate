package main

import (
	"fmt"
	"src/config"
	"src/server"
	pb "src/server/proto"

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

	job := UPDATE_USER_DATA

	switch job {
	case UPDATE_USER_DATA:
		request := &pb.UpdateUserDataRequest{}
		response, _ := grpcServer.UpdateUserData(request)
		fmt.Println(response)

	case CREATE_PLAYLIST_RECENT_IN_GENRE:
		request := &pb.CreatePlaylistRecentInGenreRequest{Genre: "Psytech"}
		response, _ := grpcServer.CreatePlaylistRecentInGenre(request)
		fmt.Println(response)

	case CREATE_PLAYLIST_RECENT_IN_GENRE_ALL:
		request := &pb.CreatePlaylistRecentInGenreAllRequest{}
		response, _ := grpcServer.CreatePlaylistRecentInGenreAll(request)
		fmt.Println(response)

	case GET_UNMAPPED_ARTISTS_FOR_USER:
		request := &pb.GetUnmappedArtistsForUserRequest{}
		response, _ := grpcServer.GetUnmappedArtistsForUser(request)
		fmt.Println(response)
	}

	if err != nil {
		fmt.Println(err)
	}
}
