package datastore_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	apptype "spotify_app/api/pkg/app_type"
	. "spotify_app/api/pkg/data_store"
	"spotify_app/api/pkg/model"
)

type ArtistTextDBTestSuite struct {
	suite.Suite
}

func TestArtistTextDBTestSuite(t *testing.T) {
	suite.Run(t, new(ArtistTextDBTestSuite))
}

func (suite *ArtistTextDBTestSuite) Test_WriteAllEntries_WritesEntries() {
	artists := []model.Artist{
		{ID: "123", Name: "HouseArtist", Genres: []apptype.Genre{"House"}},
		{ID: "456", Name: "Some Techno Guy", Genres: []apptype.Genre{"Techno"}},
	}

	db := NewArtistTextDB()

	err := db.WriteAllEntries(artists)
	suite.Require().NoError(err)
}

func (suite *ArtistTextDBTestSuite) Test_ReadArtistIDToArtist_Returns_UnmarshaledMapOfArtistIDToArtist() {
	db := NewArtistTextDB()

	artistIDToArtist, err := db.ReadArtistIDToArtist()

	suite.NotNil(artistIDToArtist)
	suite.Require().NoError(err)
}
