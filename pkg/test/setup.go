package test

import (
	"database/sql"
	"fmt"
	"order-app/pkg/docker"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewPostgresDatabase(t *testing.T) (*bun.DB, error) {
	t.Helper()

	docker.StartPgContainer(t)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"54320",
		"postgres",
		"1234",
		"postgres",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}

	var pingError error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 1000 * time.Millisecond)
	}

	err = RunMigration(db, t)
	if err != nil {
		t.Fatal(err)
	}

	b := bun.NewDB(db, pgdialect.New())
	return b, nil
}

func RunMigration(db *sql.DB, t *testing.T) error {
	t.Helper()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		t.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../migrations", "postgres", driver)
	if err != nil {
		t.Fatal(err)
	}

	return m.Up()
}
