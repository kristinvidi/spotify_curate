package datastore

import (
	"encoding/json"
	"spotify_app/api/pkg/model"
)

type ArtistTextDB struct {
	db TextDB
}

const artistFilename = "artists.json"

func NewArtistTextDB() *ArtistTextDB {
	return &ArtistTextDB{db: TextDB{}}
}

func (a *ArtistTextDB) WriteAllEntries(artists model.Artists) error {
	encodedArtists, err := artists.ArtistIDToJSONArtist()
	if err != nil {
		return err
	}

	err = a.db.WriteAllEntries(artistFilename, encodedArtists)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArtistTextDB) ReadArtistIDToArtist() (map[string]model.Artist, error) {
	artists, err := a.db.ReadAllEntries(artistFilename)
	if err != nil {
		return nil, err
	}

	var artistIDToArtist map[string]model.Artist

	err = json.Unmarshal(artists, &artistIDToArtist)
	if err != nil {
		return nil, err
	}

	return artistIDToArtist, nil
}
