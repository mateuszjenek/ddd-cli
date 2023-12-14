package port

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/mateuszjenek/ddd-cli/internal/infrastructure/port/migrations"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDatabasePort(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open an sqlite3 connection: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping an sqlite3 connection: %w", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create a migration driver: %w", err)
	}

	sourceInstance := bindata.Resource(
		migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		},
	)

	source, err := bindata.WithInstance(sourceInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to create a migration source: %w", err)
	}

	migration, err := migrate.NewWithInstance("go-bindata", source, "sqlite3", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create a migration instance: %w", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
