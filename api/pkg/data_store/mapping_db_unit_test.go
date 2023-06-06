package datastore_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	. "spotify_app/api/pkg/data_store"
)

type MappingTextDBTestSuite struct {
	suite.Suite
}

func TestMappingTextDBTestSuite(t *testing.T) {
	suite.Run(t, new(MappingTextDBTestSuite))
}

func (suite *MappingTextDBTestSuite) Test_ReadArtistIDToArtist_Returns_UnmarshaledMapOfArtistIDToArtist() {
	db := NewMappingTextDB()

	genreToArtists, err := db.ReadGenreToArtistsMapping()
	fmt.Println(genreToArtists)

	suite.NotNil(genreToArtists)
	suite.Require().NoError(err)
}
