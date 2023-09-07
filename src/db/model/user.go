package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user"`

	ID          string `bun:"spotify_id"`
	URI         string `bun:"uri"`
	DisplayName string `bun:"display_name"`
	Email       string `bun:"email"`
	Country     string `bun:"country"`
}
