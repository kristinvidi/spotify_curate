package domain

import (
	"src/config"
	"src/db/connection"
	"src/db/query"
	mapperapi "src/domain/mapper_api"
	mapperdb "src/domain/mapper_db"
	"src/spotifyapi"
)

type User struct {
	config *config.Config
}

func NewUser(config *config.Config) *User {
	return &User{config: config}
}

func (u *User) GetAndStoreCurrentUsersProfile() error {
	api := spotifyapi.GetUser(u.config)
	profile, err := api.GetCurrentUsersProfile()
	if err != nil {
		return err
	}

	user := mapperdb.UserToDBUser(
		mapperapi.UserFromCurrentUsersProfileResponse(profile),
	)

	dbConnection := connection.GetConnection(u.config.Database)
	dbUser := query.NewUser(dbConnection)

	err = dbUser.InsertUserData(*user)
	if err != nil {
		return err
	}

	return nil
}
