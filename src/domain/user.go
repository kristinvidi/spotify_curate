package domain

import (
	"src/config"
	"src/db"
	"src/domain/mapper"
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
	response, err := api.GetCurrentUsersProfile()
	if err != nil {
		return err
	}

	user := mapper.DBUserFromCurrentUsersProfileResponse(response)

	dbUser := db.GetUser(u.config.Database)

	err = dbUser.InsertUserData(*user)
	if err != nil {
		return err
	}

	return nil
}
