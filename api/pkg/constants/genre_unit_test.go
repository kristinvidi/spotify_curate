package constants_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	. "spotify_app/api/pkg/constants"
)

type GenreTestSuite struct {
	suite.Suite
}

func TestGenreTestSuite(t *testing.T) {
	suite.Run(t, new(GenreTestSuite))
}

func (suite *GenreTestSuite) Test_String_Returns_StringForEachConstant() {
	suite.Equal(GenreDeepHouse.String(), "Deep House")
}

func (suite *GenreTestSuite) Test_GenreFromString_Returns_Genre_When_StringIsValidGenre() {
	g, _ := GenreFromString("Deep House")

	suite.Equal(*g, GenreDeepHouse)
}
