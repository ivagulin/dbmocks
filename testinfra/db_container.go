package testinfra

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"

	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var _ DB = (*containerDB)(nil)

type containerDB struct {
	ctr     *postgres.PostgresContainer
	connStr string
}

func newContainerDB() (db *containerDB, err error) {
	// run container

	ctx := context.Background()
	ctr, err := postgres.Run(ctx, "postgres:15",
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2)),
		postgres.WithDatabase("integration_tests_db"),
		testcontainers.CustomizeRequestOption(func(req *testcontainers.GenericContainerRequest) error {
			req.Cmd = append(req.Cmd, "-c", "log_statement=all")
			return nil
		}),
	)
	defer func() {
		if err != nil {
			_ = testcontainers.TerminateContainer(ctr)
		}
	}()
	if err != nil {
		return nil, err
	}

	connStr, err := ctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("get postgres connection string: %s", err)
	}
	pConnCfg, err := pgx.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parse postgres connection string: %s", err)
	}
	if err := applyMigrations(*pConnCfg); err != nil {
		return nil, fmt.Errorf("apply migrations: %s", err)
	}

	// take a snapshot that can be restored

	if err := ctr.Snapshot(ctx); err != nil {
		return nil, fmt.Errorf("take db snapshot: %s", err)
	}

	return &containerDB{
		ctr:     ctr,
		connStr: connStr,
	}, nil
}

func (db *containerDB) ConnStr() string {
	return db.connStr
}

func (db *containerDB) MustCleanup() {
	lo.Must0(db.ctr.Restore(context.Background()))
}

func (db *containerDB) Close() {
	_ = db.ctr.Terminate(context.Background())
}
