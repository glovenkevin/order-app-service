package app

import (
	"database/sql"
	"errors"
	"order-app/pkg/logger"
	"time"

	"github.com/golang-migrate/migrate/v4"

	// migrate tools.
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func execMigration(db *sql.DB, log logger.LoggerInterface) error {
	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	for attempts > 0 {
		m, err = migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
		if err == nil {
			break
		}

		log.Warnf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
		return err
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
		return err
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info("Migrate: no change")
		return nil
	}

	log.Info("Migrate: up success")
	return nil
}

func execMigrationDown(db *sql.DB, log logger.LoggerInterface) error {
	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	for attempts > 0 {
		m, err = migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
		if err == nil {
			break
		}

		log.Warnf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
		return err
	}

	err = m.Down()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: down error: %s", err)
		return err
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info("Migrate: no change")
		return nil
	}

	log.Info("Migrate: down success")
	return nil
}
