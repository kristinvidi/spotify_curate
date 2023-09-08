package query

import (
	"src/db/connection"
	"src/db/model"
)

type User struct {
	db *connection.PostgresDB
}

func NewUser(db *connection.PostgresDB) *User {
	return &User{db: db}
}

func (u *User) InsertUserData(user model.User) error {
	return u.db.Insert(&user)
}
