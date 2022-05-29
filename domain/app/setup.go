package app

import (
	"database/sql"
	"fmt"
	"order-app/config"
	"order-app/pkg/httpserver"
	"order-app/pkg/logger"
	"order-app/pkg/postgres"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-pg/pg/v10"
)

func initLogger(cfg *config.Log) (logger.ILogger, error) {
	log, err := logger.NewZapLogger(cfg)
	if err != nil {
		return nil, err
	}
	return log, nil
}

func initDatabase(cfg *config.PG) (*pg.DB, error) {
	db, err := postgres.NewORM(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.PoolMax)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initDatabaseMigration(cfg *config.PG, log logger.ILogger) (*sql.DB, error) {
	db, err := postgres.New(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.PoolMax, log)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func catchSignal(httpServer *httpserver.Server, log logger.ILogger) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
