package query

import (
	"fmt"

	"context"
	"src/config"
	"src/db/connection"
	"src/db/constants"

	"github.com/uptrace/bun"
)

type PostgresDB struct {
	db bun.IDB
}

func NewPostgresDB(dbConfig config.DB) *PostgresDB {
	db := connection.GetConnection(dbConfig)

	return &PostgresDB{db: db}
}

func (p *PostgresDB) insert(m interface{}, conflictConstraint constants.Column, conflictCommand constants.OnConflict) error {
	conflictClause := fmt.Sprintf("CONFLICT (%s) %s", conflictConstraint, conflictCommand)
	_, err := p.db.NewInsert().Model(m).On(conflictClause).Exec(context.Background())

	return err
}
