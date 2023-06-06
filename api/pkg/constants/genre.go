package constants

import "fmt"

type Genre int32

const (
	GenreChunkyTechHouse Genre = iota
	GenreClassicHouse
	GenreDeepAndProgressiveHouse
	GenreDiscoAndJazzyHouse
	GenreDrumAndBass
	GenreDubstep
	GenreHouse
	GenreMinimalTechno
	GenreProgressiveHouseAndTechno
	GenreProgressiveTechHouse
	GenrePsytech
	GenreTechHouse
)

func genreToString() map[Genre]string {
	genreToString := make(map[Genre]string)
	genreToString[GenreChunkyTechHouse] = "Chunky Tech House"
	genreToString[GenreClassicHouse] = "Classic House"
	genreToString[GenreDeepAndProgressiveHouse] = "Deep and Progressive House"
	genreToString[GenreDiscoAndJazzyHouse] = "Disco and Jazzy House"
	genreToString[GenreDrumAndBass] = "Drum and Bass"
	genreToString[GenreDubstep] = "Dubstep"
	genreToString[GenreHouse] = "House"
	genreToString[GenreMinimalTechno] = "Minimal Techno"
	genreToString[GenreProgressiveHouseAndTechno] = "Progressive House and Techno"
	genreToString[GenreProgressiveTechHouse] = "Progressive Tech House"
	genreToString[GenrePsytech] = "Psytech"
	genreToString[GenreTechHouse] = "Tech House"

	return genreToString
}

func stringToGenre() map[string]Genre {
	stringToGenre := make(map[string]Genre)

	for g, s := range genreToString() {
		stringToGenre[s] = g
	}

	return stringToGenre
}

func (g Genre) String() string {
	genreToString := genreToString()

	return genreToString[g]
}

func GenreFromString(s string) (*Genre, error) {
	stringToGenre := stringToGenre()

	g, ok := stringToGenre[s]
	if !ok {
		return nil, fmt.Errorf("could not find genre %s", s)
	}

	return &g, nil
}
