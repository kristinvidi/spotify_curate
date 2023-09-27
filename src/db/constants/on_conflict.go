package constants

type OnConflict string

const (
	OnConflictDoNothing OnConflict = "DO NOTHING"
	OnConflictDoUpdate  OnConflict = "DO UPDATE"
)
