package query

import (
	"context"
	"src/db/model"

	"github.com/uptrace/bun"
)

type User struct {
	db *bun.DB
}

func NewUser(db *bun.DB) *User {
	return &User{db: db}
}

func (u *User) InsertUserData(user model.User) error {
	ctx := context.Background()
	_, err := u.db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUser(id int64) (*model.User, error) {
	ctx := context.Background()
	var user model.User
	err := u.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
