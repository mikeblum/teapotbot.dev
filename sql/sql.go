package sql

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/knadh/koanf"
	"github.com/mikeblum/teapotbot.dev/conf"
	"github.com/pressly/goose/v3"
)

const (
	dialect         = "postgres"
	driver          = "pgx"
	migrationsDir   = "internal/migrations"
	migrationsTable = "migrations"
)

//go:embed internal/migrations/*.sql
var embedMigrations embed.FS

func Setup(cfg *koanf.Koanf) error {
	var db *sql.DB
	var err error
	if db, err = sql.Open(driver, cfg.String(EnvDatabaseUrl)); err != nil {
		return err
	}
	if db == nil {
		return fmt.Errorf("db not configured")
	}

	defer db.Close()

	goose.SetBaseFS(embedMigrations)
	goose.SetTableName(migrationsTable)
	goose.SetLogger(conf.NewLog(""))
	goose.SetSequential(true)

	if err = goose.SetDialect(dialect); err != nil {
		return err
	}
	// TODO: remove this
	if err = goose.Reset(db, migrationsDir); err != nil {
		return nil
	}
	if err = goose.Up(db, migrationsDir); err != nil {
		return err
	}
	return nil
}
