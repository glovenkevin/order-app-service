package test

import (
	"log"
	"order-app/config"
	v1 "order-app/domain/controller/http/v1"
	"order-app/pkg/httpserver"
	"order-app/pkg/logger"
	"order-app/pkg/postgres"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/uptrace/bun"
)

type TestSuite struct {
	h  *httpserver.Server
	db *bun.DB
}

func (ts *TestSuite) TearDown() {
	ts.h.Shutdown()
	ts.db.Close()
}

func (ts *TestSuite) ExecQuery(qs []string) {
	for _, q := range qs {
		_, err := ts.db.Exec(q)
		if err != nil {
			log.Fatalf("Exec query error: %s", err)
		}
	}
}

func loadConfig() *config.Config {
	cfg := &config.Config{}

	err := cleanenv.ReadConfig("config.test.yml", cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	return cfg
}

func initDb(cfg *config.PG, log logger.LoggerInterface) (*bun.DB, error) {
	db, err := postgres.NewORM(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.PoolMax, log)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newLogger(cfg *config.Log) (logger.LoggerInterface, error) {
	log, err := logger.NewZapLogger(cfg)
	if err != nil {
		return nil, err
	}
	return log, nil
}

func NewTestSuite() *TestSuite {
	cfg := loadConfig()
	l, err := newLogger(&cfg.Log)
	db, err := initDb(&cfg.PG, l)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	handler := httpserver.NewServerHandler(&cfg.App, l)
	v1.NewRouter(handler, l, db)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	return &TestSuite{h: httpServer, db: db}
}
