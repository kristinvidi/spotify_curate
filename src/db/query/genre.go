package query

import (
	"context"
	"database/sql"
	"fmt"

	"src/db/model"
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
