package data_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/suite"

// 	apptype "spotify_app/api/pkg/app_type"
// 	. "spotify_app/api/pkg/data"
// 	"spotify_app/api/pkg/model"
// )

// type TopGenresTestSuite struct {
// 	suite.Suite
// }

// func TestTopGenresTestSuite(t *testing.T) {
// 	suite.Run(t, new(TopGenresTestSuite))
// }

// func buildArtist(name string, genreList apptype.Genres) model.Artist {
// 	return model.Artist{Name: name, Genres: genreList}
// }

// func (suite *TopGenresTestSuite) Test_GetMapOfGenreToArtists_Returns_Map_WithSliceOfArtistsByGenre_When_MultipleArtistsInSlice_With_SliceOfGenresNotEmpty() {
// 	artistList := []model.Artist{
// 		buildArtist("A Cool Artist", apptype.Genres{apptype.GenreDisco, apptype.GenreHouse}),
// 		buildArtist("A Cooler Artist", apptype.Genres{apptype.GenreHouse, apptype.GenreTechno}),
// 		buildArtist("No Genres", apptype.Genres{}),
// 		buildArtist("No Genres Again", apptype.Genres{}),
// 	}

// 	mapOfGenreToArtists, artistsWithoutGenres := GetMapOfGenreToArtists(artistList)

// 	suite.Equal([]string{"A Cool Artist"}, mapOfGenreToArtists[apptype.GenreDisco])
// 	suite.Equal([]string{"A Cool Artist", "A Cooler Artist"}, mapOfGenreToArtists[apptype.GenreHouse])
// 	suite.Equal([]string{"A Cooler Artist"}, mapOfGenreToArtists[apptype.GenreTechno])
// 	suite.Equal("No Genres", artistsWithoutGenres[0].Name)
// 	suite.Equal("No Genres Again", artistsWithoutGenres[1].Name)
// }
