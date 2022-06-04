package app

import (
	"context"
	"fmt"
	"log"
	"order-app/config"
	"order-app/pkg/httpserver"
	"os"

	v1 "order-app/domain/controller/http/v1"

	"github.com/urfave/cli/v2"
)

func Run(cfg *config.Config) {
	app := cli.NewApp()
	app.Name = cfg.App.Name
	app.Version = cfg.App.Version
	app.Action = runCommand(cfg)
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "mode",
			Required: false,
			Value:    "app",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(cfg *config.Config) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		mode := c.String("mode")
		switch mode {
		case "migration":
			return runMigrate(cfg)
		case "app":
			return runApp(cfg)
		default:
			return fmt.Errorf("app - Run mode unrecognize: %s", mode)
		}
	}
}

func runApp(cfg *config.Config) error {
	l, err := initLogger(&cfg.Log)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer l.Close()

	db, err := initDatabase(&cfg.PG)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
		return err
	}
	defer db.Close()

	// HTTP Server
	handler := httpserver.NewServerHandler(&cfg.App)
	v1.NewRouter(handler, l, db)

	// Start the server and wait for the interrupt signal to gracefully shutdown the server with
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Infof("App running on port %s", cfg.HTTP.Port)
	catchSignal(httpServer, l)

	return nil
}

func runMigrate(cfg *config.Config) error {
	l, err := initLogger(&cfg.Log)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer l.Close()

	db, err := initDatabaseMigration(&cfg.PG, l)
	if err != nil {
		l.Fatalf("app - Run - postgres.New: %w", err)
		return err
	}

	// Migrate
	ctx := context.Background()
	return execMigration(ctx, db, l)
}