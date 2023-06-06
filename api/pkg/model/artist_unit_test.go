package model_test

import (
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/model"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TopGenresTestSuite struct {
	suite.Suite
}

func TestTopGenresTestSuite(t *testing.T) {
	suite.Run(t, new(TopGenresTestSuite))
}

func (suite *TopGenresTestSuite) Test_Artists_TopGenre_Returns_FirstGenre_When_ArtistsIsOneArtist() {
	artists := model.Artists{
		{Name: "", Genres: []apptype.Genre{"House", "Disco", "Techno"}},
	}

	topGenre := artists.TopGenre()

	suite.Equal(apptype.Genre("House"), *topGenre)
}

func (suite *TopGenresTestSuite) Test_Artists_TopGenre_Returns_FirstGenreWithHighestCount_When_MultipleGenresHaveHighestCount() {
	artists := model.Artists{
		{Name: "", Genres: []apptype.Genre{"House", "Disco", "Techno"}},
		{Name: "", Genres: []apptype.Genre{"House", "DnB", "Disco"}},
	}

	topGenre := artists.TopGenre()

	suite.Equal(apptype.Genre("House"), *topGenre)
}

func (suite *TopGenresTestSuite) Test_Artists_TopGenre_Returns_GenreWithHighestCount_When_MultipleArtists() {
	artists := model.Artists{
		{Name: "", Genres: []apptype.Genre{"House", "Disco", "Techno"}},
		{Name: "", Genres: []apptype.Genre{"House", "DnB", "Disco"}},
		{Name: "", Genres: []apptype.Genre{"Party", "DnB", "Disco"}},
	}

	topGenre := artists.TopGenre()

	suite.Equal(apptype.Genre("Disco"), *topGenre)
}

func (suite *TopGenresTestSuite) Test_Artists_TopGenre_Returns_GenreWithHighestCount_When_SingleArtistHasMultipleOfSameGenre() {
	artists := model.Artists{
		{Name: "", Genres: []apptype.Genre{"House", "Disco", "Techno", "Disco"}},
	}

	topGenre := artists.TopGenre()

	suite.Equal(apptype.Genre("Disco"), *topGenre)
}

func (suite *TopGenresTestSuite) Test_Artists_TopGenre_Returns_GenreWithHighestCount_When_SingleArtistHasMultipleOfSameGenre_And_SecondArtistHasNoGenres() {
	artists := model.Artists{
		{Name: "", Genres: []apptype.Genre{"House", "Disco", "Techno", "Disco"}},
		{},
	}

	topGenre := artists.TopGenre()

	suite.Equal(apptype.Genre("Disco"), *topGenre)
}

func (suite *TopGenresTestSuite) Test_Artists_TopGenre_Returns_Nil_When_ArtistsHaveNoGenres() {
	artists := model.Artists{
		{},
		{},
	}

	topGenre := artists.TopGenre()

	suite.Nil(topGenre)
}

func (suite *TopGenresTestSuite) Test_Artists_TopGenre_Returns_Nil_When_ArtistsIsEmpty() {
	artists := model.Artists{}

	topGenre := artists.TopGenre()

	suite.Nil(topGenre)
}
