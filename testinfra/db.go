package testinfra

import (
	"github.com/ivagulin/dbmocks/repo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/lo"
	"os"
)

const ENV_NAME_DB_CONN_STR = "TEST_DB_CONN_STR"

// DB represents either a local database container or an externally provided database
// (via the ENV_NAME_DB_CONN_STR environment variable; e.g., it can be a GitLab CI/CD service).
type DB interface {
	ConnStr() string
	// Attention: all connection must be restarted
	MustCleanup()
	Close()
}

func MustStartDB() DB {
	return lo.Must(StartDB())
}

func StartDB() (DB, error) {
	connStr := os.Getenv(ENV_NAME_DB_CONN_STR)
	if connStr != "" {
		return newExternalDB(connStr)
	}
	return newContainerDB()
}

func applyMigrations(connCfg pgx.ConnConfig) error {
	conn := stdlib.OpenDB(connCfg)
	defer conn.Close()
	return repo.MigrateDB(conn)
}
