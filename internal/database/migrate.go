package database

import (
	stdsql "database/sql"
	"fmt"
	"io"
	"log"

	"github.com/MontillaTomas/blog-aggregator/sql"
	"github.com/pressly/goose/v3"
)

func Migrate(db *stdsql.DB) error {
	goose.SetBaseFS(sql.Migrations)
	goose.SetLogger(log.New(io.Discard, "", 0))

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	if err := goose.Up(db, "schema"); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
