package connection

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"src/config"
)

func GetConnection(dbConfig config.DB) *bun.DB {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"

	connection := pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
		pgdriver.WithDatabase(dbConfig.Database),
		pgdriver.WithUser(dbConfig.User),
		pgdriver.WithPassword(dbConfig.Password),
		pgdriver.WithTimeout(5*time.Second),
	)

	db := bun.NewDB(sql.OpenDB(connection), pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return db
}
