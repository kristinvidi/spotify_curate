package datastore

import (
	"encoding/json"
	"spotify_app/api/pkg/constants"
	"strings"
)

type MappingTextDB struct {
	db TextDB
}

const mappingFilename = "artist_genre_mapping.json"

func NewMappingTextDB() *MappingTextDB {
	return &MappingTextDB{db: TextDB{}}
}

func (a *MappingTextDB) ReadGenreToArtistsMapping() (map[constants.Genre][]string, error) {
	store, err := a.db.ReadAllEntries(mappingFilename)
	if err != nil {
		return nil, err
	}

	genreToArtists := make(map[string][]string)

	err = json.Unmarshal(store, &genreToArtists)
	if err != nil {
		return nil, err
	}

	genreToArtistsClean := make(map[constants.Genre][]string)
	for g, artists := range genreToArtists {
		genre, err := constants.GenreFromString(g)
		if err != nil {
			return nil, err
		}

		genreToArtistsClean[*genre] = splitArtistByHash(artists)
	}

	return genreToArtistsClean, nil
}

func splitArtistByHash(artistHash []string) []string {
	var splitArtists []string
	for _, a := range artistHash {
		split := strings.Split(a, "#")
		splitArtists = append(splitArtists, split[0])
	}

	return splitArtists
}
