package model

import "time"

type Artist struct {
	ID        ID
	URI       URI
	Name      string
	CreatedAt time.Time
}
