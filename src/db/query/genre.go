package query

import (
	"context"
	"database/sql"
	"fmt"

	"src/db/constants"
	"src/db/model"

	"github.com/uptrace/bun"
)

func (p *PostgresDB) GetGenreMappingForUserAndGenre(userID model.ID, genre model.Genre) (*model.UserIDGenreMapping, error) {
	var genreMapping model.UserIDGenreMapping

	err := p.db.NewSelect().
		Model(&genreMapping).
		Where("user_spotify_id = ?", userID).
		Where("genre = ?", genre).
		Scan(context.Background())

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("genre %s is not mapped for user", genre)
	}

	if err != nil {
		return nil, err
	}

	return &genreMapping, err
}

func (p *PostgresDB) GetGenreMappingsForUser(userID model.ID) ([]model.UserIDGenreMapping, error) {
	var genreMappings []model.UserIDGenreMapping

	err := p.db.NewSelect().
		Model(&genreMappings).
		Where("user_spotify_id = ?", userID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return genreMappings, nil
}

func (p *PostgresDB) GetGenreMappingForUserAndGenreName(userID model.ID, genreName string) (*model.UserIDGenreMapping, error) {
	var genreMapping model.UserIDGenreMapping

	err := p.db.NewSelect().
		Model(&genreMapping).
		Where("? = ?", bun.Ident(constants.ColumnGenre.String()), genreName).
		Scan(context.Background())

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("genre %s is not mapped for user", genreName)
	}

	if err != nil {
		return nil, err
	}

	return &genreMapping, err
}

func (p *PostgresDB) InsertUserArtistIDGenreMappings(mappings model.UserArtistIDGenreMappings) error {
	if len(mappings) == 0 {
		return nil
	}

	return p.insertWithConflict(&mappings, constants.ColumnUserArtistGenreID, constants.OnConflictDoNothing)
}
