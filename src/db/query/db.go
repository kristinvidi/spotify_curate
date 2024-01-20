package query

import (
	"fmt"

	"context"
	"src/config"
	"src/db/connection"
	"src/db/constants"
	"src/db/model"

	"github.com/uptrace/bun"
)

type PostgresDB struct {
	db bun.IDB
}

func NewPostgresDB(dbConfig config.DB) *PostgresDB {
	db := connection.GetConnection(dbConfig)

	return &PostgresDB{db: db}
}

func (p *PostgresDB) insertWithConflict(m interface{}, conflictConstraint constants.Column, conflictCommand constants.OnConflict) error {
	conflictClause := fmt.Sprintf("CONFLICT (%s) %s", conflictConstraint, conflictCommand)
	_, err := p.db.NewInsert().Model(m).On(conflictClause).Exec(context.Background())

	return err
}

func (p *PostgresDB) insertNoConflict(m interface{}) error {
	_, err := p.db.NewInsert().Model(m).Exec(context.Background())

	return err
}

func (p *PostgresDB) deleteByUserInAndArtistIDNotIn(m interface{}, userID model.ID, artistIDs []model.ID) (int64, error) {
	res, err := p.db.NewDelete().
		Model(m).
		Where("? = ?", bun.Ident(constants.ColumnUserID), userID).
		Where("? not in (?)", bun.Ident(constants.ColumnArtistID), bun.In(artistIDs)).
		Exec(context.Background())

	rows, _ := res.RowsAffected()

	return rows, err
}
