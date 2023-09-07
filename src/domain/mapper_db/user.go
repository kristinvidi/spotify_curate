package mapperdb

import (
	dbModel "src/db/model"
	"src/domain/model"
)

func UserToDBUser(user *model.User) *dbModel.User {
	if user == nil {
		return nil
	}

	return &dbModel.User{
		DisplayName: string(user.DisplayName),
		Email:       string(user.Email),
		ID:          string(user.ID),
		URI:         string(user.URI),
		Country:     string(user.Country),
	}
}
