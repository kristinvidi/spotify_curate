package main

import (
	"src/config"
	"src/server"

	"go.uber.org/zap"
)

func main() {
	// Get config / environment variables
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	// Set up logger
	logger := zap.Must(zap.NewDevelopment())
	if config.AppEnv.Env == "production" {
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync()

	// Set up and start server
	grpcServer := server.NewGrpcServer(config, logger)
	if err := grpcServer.Run(); err != nil {
		panic(err)
	}
}
