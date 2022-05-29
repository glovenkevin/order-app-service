package postgres

import (
	"database/sql"
	"fmt"

	"order-app/pkg/logger"

	"github.com/go-pg/pg/v10"
)

func New(host, port, user, password, database string, pool int, log logger.ILogger) (*sql.DB, error) {
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

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewORM(host, port, user, password, database string, pool int) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		User:     user,
		Password: password,
		Database: database,
		PoolSize: pool,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
