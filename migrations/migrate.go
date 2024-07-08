package migrations

import (
	"errors"
	"fmt"
	"os"
	"product-storage/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationPathTemplate = "file://%s/migrations"

func MigrateUp(cfg config.DBConfig) error {
	connStr := fmt.Sprintf(
		"pgx5://%s:%s@%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Addr, cfg.Name, cfg.SSLMode,
	)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	migrationPath := fmt.Sprintf(migrationPathTemplate, wd)
	m, err := migrate.New(migrationPath, connStr)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
