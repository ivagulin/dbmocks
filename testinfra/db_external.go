package testinfra

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/lo"
)

var _ DB = (*externalDB)(nil)

type externalDB struct {
	connCfg      pgx.ConnConfig
	connStr      string
	snapshotName string
}

// newExternalDB creates a new DB instance that connects to an existing database (e.g. postgres service provided by gitlab CI).
// To enable concurrent integration tests, a new database with a random name is created on start and dropped on close.
//
//	Alternative solution: run integration tests sequentially, create snapshot at first newExternalDB call.
func newExternalDB(connStr string) (*externalDB, error) {
	pConnCfg, err := pgx.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parse connection string: %s", err)
	}
	connCfg := *pConnCfg

	connCfg, err = createNewDatabase(connCfg)
	if err != nil {
		return nil, fmt.Errorf("create new database: %s", err)
	}
	parsedURL, _ := url.Parse(connStr)
	parsedURL.Path = connCfg.Database
	connStr = parsedURL.String()

	if err := applyMigrations(connCfg); err != nil {
		return nil, fmt.Errorf("apply migrations: %s", err)
	}

	db := &externalDB{
		connCfg:      connCfg,
		connStr:      connStr,
		snapshotName: connCfg.Database + "_snapshot",
	}
	if err := db.takeSnapshot(); err != nil {
		db.Close()
		return nil, fmt.Errorf("take db snapshot: %s", err)
	}

	return db, nil
}

func (db *externalDB) ConnStr() string {
	return db.connStr
}

func (db *externalDB) MustCleanup() {
	lo.Must0(db.restoreSnapshot())
}

func (db *externalDB) Close() {
	lo.Must0(dropDatabase(db.connCfg))
}

func (db *externalDB) takeSnapshot() error {
	conn := stdlib.OpenDB(db.connCfg)
	defer func() { _ = conn.Close() }()

	// copy from github.com/testcontainers/testcontainers-go/modules/postgres
	cmds := []string{
		// Update pg_database to remove the template flag, then drop the database if it exists.
		// This is needed because dropping a template database will fail.
		// https://www.postgresql.org/docs/current/manage-ag-templatedbs.html
		fmt.Sprintf(`UPDATE pg_database SET datistemplate = FALSE WHERE datname = '%s'`, db.snapshotName),
		fmt.Sprintf(`DROP DATABASE IF EXISTS "%s"`, db.snapshotName),
		// Create a copy of the database to another database to use as a template now that it was fully migrated
		fmt.Sprintf(`CREATE DATABASE "%s" WITH TEMPLATE "%s" OWNER "%s"`, db.snapshotName, db.connCfg.Database, db.connCfg.User),
		// Snapshot the template database so we can restore it onto our original database going forward
		fmt.Sprintf(`ALTER DATABASE "%s" WITH is_template = TRUE`, db.snapshotName),
	}
	for _, cmd := range cmds {
		if _, err := conn.Exec(cmd); err != nil {
			return fmt.Errorf("execute command %s: %s", cmd, err)
		}
	}

	return nil
}

func (db *externalDB) restoreSnapshot() error {
	// connect to 'postgres' db to drop our target database
	connCfg := db.connCfg
	connCfg.Database = "postgres"
	connPool := stdlib.OpenDB(connCfg)
	defer func() { _ = connPool.Close() }()

	// copy from github.com/testcontainers/testcontainers-go/modules/postgres
	cmds := []string{
		// Drop the entire database by connecting to the postgres global database
		fmt.Sprintf(`DROP DATABASE "%s" with (FORCE)`, db.connCfg.Database),
		// Then restore the previous snapshot
		fmt.Sprintf(`CREATE DATABASE "%s" WITH TEMPLATE "%s" OWNER "%s"`, db.connCfg.Database, db.snapshotName, db.connCfg.User),
	}
	for _, cmd := range cmds {
		if _, err := connPool.Exec(cmd); err != nil {
			return fmt.Errorf("execute command %s: %s", cmd, err)
		}
	}

	return nil
}

func createNewDatabase(connCfg pgx.ConnConfig) (pgx.ConnConfig, error) {
	conn := stdlib.OpenDB(connCfg)
	defer func() { _ = conn.Close() }()

	newDatabseName := fmt.Sprintf("%s_%x", connCfg.Database, rand.Int())
	if _, err := conn.Exec("CREATE DATABASE " + newDatabseName); err != nil {
		return connCfg, err
	}
	connCfg.Database = newDatabseName

	return connCfg, nil
}

func dropDatabase(connCfg pgx.ConnConfig) error {
	dbName := connCfg.Database
	// connect to postgres db to drop target db.
	connCfg.Database = "postgres"
	conn := stdlib.OpenDB(connCfg)
	defer func() { _ = conn.Close() }()

	_, err := conn.Exec("DROP DATABASE " + dbName)
	return err
}
