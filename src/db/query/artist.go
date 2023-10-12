package query

import (
	"src/db/constants"
	"src/db/model"
)

func (p *PostgresDB) InsertArtistData(artists []model.Artist) error {
	return p.insertWithConflict(&artists, constants.ColumnID, constants.OnConflictDoNothing)
}
