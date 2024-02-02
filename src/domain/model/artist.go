package model

import (
	"time"
)

type Artist struct {
	ID        ID
	URI       URI
	Name      string
	CreatedAt time.Time
}

type Artists []Artist

func (a Artists) Names() []string {
	names := make([]string, len(a))
	for _, artist := range a {
		names = append(names, artist.Name)
	}

	return names
}
