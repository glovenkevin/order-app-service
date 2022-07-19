package postgres

import (
	"database/sql"
	"fmt"

	"order-app/pkg/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func New(host, port, user, password, database string, pool int, log logger.LoggerInterface) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		database,
	)

	log.Info(fmt.Sprintf(
		"host=%s port=%s dbname=%s",
		host,
		port,
		database,
	))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if pool != 0 {
		db.SetMaxOpenConns(pool)
	}

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewORM(host, port, user, password, database string, pool int, log logger.LoggerInterface) (*bun.DB, error) {
	conn, err := New(host, port, user, password, database, pool, log)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(conn, pgdialect.New())

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
