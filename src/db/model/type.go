package model

type (
	ID    string
	URI   string
	Genre string
)

type ReleaseDatePrecision string

const (
	ReleaseDatePrecisionYear  ReleaseDatePrecision = "year"
	ReleaseDatePrecisionMonth ReleaseDatePrecision = "month"
	ReleaseDatePrecisionDay   ReleaseDatePrecision = "day"
)
