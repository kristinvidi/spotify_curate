package main

import (
	"fmt"
	"src/config"
	"src/server"
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

	grpcServer := server.NewGrpcServer(config, nil)

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

// func createFormattedFieldLogger(fields logrus.Fields) *logrus.Logger {
// 	logger := logrus.Logger{
// 		Out:   os.Stderr,
// 		Level: logrus.GetLevel(),
// 		Formatter: &easy.Formatter{
// 			LogFormat:       "%thread% %time% [%lvl%] %msg%",
// 			TimestampFormat: logTimeFormat,
// 		},
// 	}

// 	fieldLogger := logger.WithFields(fields)

// 	return fieldLogger
// }
