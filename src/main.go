package main

import (
	"src/config"
	"src/domain/update"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	user := update.NewUserData(config)
	err = user.UpdateAllUserData()
	if err != nil {
		panic(err)
	}
}
