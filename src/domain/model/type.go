package model

type ID string
type URI string
type Genre string

type AlbumType string

const (
	AlbumTypeAlbum       AlbumType = "album"
	AlbumTypeSingle      AlbumType = "single"
	AlbumTypeCompilation AlbumType = "compilation"
)
