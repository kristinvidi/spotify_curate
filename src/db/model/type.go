package model

type ID string
type URI string
type Genre string

type ReleaseDatePrecision string

const (
	ReleaseDatePrecisionYear  ReleaseDatePrecision = "year"
	ReleaseDatePrecisionMonth ReleaseDatePrecision = "month"
	ReleaseDatePrecisionDay   ReleaseDatePrecision = "day"
)
