package connection

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"src/config"
)

type PostgresDB struct {
	db     bun.IDB
	logger logrus.FieldLogger
}

func NewPostgresDB(db bun.IDB, logger logrus.FieldLogger) *PostgresDB {
	return &PostgresDB{db: db, logger: logger}
}

func GetConnection(dbConfig config.DB) *bun.DB {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"

	connection := pgdriver.NewConnector(
		// pgdriver.WithAddr(dbConfig.Host+":"+dbConfig.Port),
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
