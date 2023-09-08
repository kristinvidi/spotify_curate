package db

import (
	"src/config"
	"src/db/connection"
	"src/db/query"

	"github.com/sirupsen/logrus"
)

func GetUser(dbConfig config.DB) *query.User {
	postgresDB := connection.NewPostgresDB(dbConfig, logrus.New())
	return query.NewUser(postgresDB)
}
