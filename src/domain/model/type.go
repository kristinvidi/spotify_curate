package model

type (
	ID    string
	URI   string
	Genre string
)

type AlbumType string

const (
	AlbumTypeAlbum       AlbumType = "album"
	AlbumTypeSingle      AlbumType = "single"
	AlbumTypeCompilation AlbumType = "compilation"
)
