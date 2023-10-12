package model

type ID string

func (s ID) String() string {
	return string(s)
}

type URI string

type AlbumType string

const (
	AlbumTypeAlbum       AlbumType = "album"
	AlbumTypeSingle      AlbumType = "single"
	AlbumTypeCompilation AlbumType = "compilation"
)

type ReleaseDatePrecision string

const (
	ReleaseDatePrecisionYear  ReleaseDatePrecision = "year"
	ReleaseDatePrecisionMonth ReleaseDatePrecision = "month"
	ReleaseDatePrecisionDay   ReleaseDatePrecision = "day"
)
