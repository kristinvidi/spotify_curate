package main

import (
	"src/config"
	"src/domain"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	user := domain.NewUser(config)
	err = user.GetAndStoreCurrentUsersProfile()
	if err != nil {
		panic(err)
	}
}
