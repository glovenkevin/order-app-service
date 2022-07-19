package test

import (
	"database/sql"
	"order-app/config"
	"order-app/pkg/docker"
	"order-app/pkg/logger"
	"testing"

	"github.com/uptrace/bun"
)

type TestCase struct {
	Db       *bun.DB
	Sql      *sql.DB
	TearDown func()
	Log      logger.LoggerInterface
}

func NewTestCase(t *testing.T) *TestCase {
	b, err := NewPostgresDatabase(t)
	if err != nil {
		t.Fatal(err)
	}

	log, err := logger.NewZapLogger(&config.Log{
		Level: "debug",
	})

	return &TestCase{
		Db:  b,
		Sql: b.DB,
		Log: log,
		TearDown: func() {
			t.Helper()
			b.Close()
			docker.StopPgContainer(t)
		},
	}
}
