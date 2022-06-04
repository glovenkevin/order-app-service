package test

import (
	"log"
	"order-app/config"
	v1 "order-app/domain/controller/http/v1"
	"order-app/pkg/httpserver"
	"order-app/pkg/logger"
	"order-app/pkg/postgres"

	"github.com/go-pg/pg/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type TestSuite struct {
	h  *httpserver.Server
	db *pg.DB
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

func initDb(cfg *config.PG) (*pg.DB, error) {
	db, err := postgres.NewORM(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.PoolMax)
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
	db, err := initDb(&cfg.PG)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	l, err := newLogger(&cfg.Log)

	handler := httpserver.NewServerHandler(&cfg.App)
	v1.NewRouter(handler, l, db)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	return &TestSuite{h: httpServer, db: db}
}
