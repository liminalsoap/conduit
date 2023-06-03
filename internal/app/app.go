package app

import (
	"conduit-go/config"
	"conduit-go/pkg/logger"
	"conduit-go/pkg/postgres"
	"context"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger(cfg.Logger.Level)

	conn, err := postgres.NewDb(cfg.Postgresql.Url)
	if err != nil {
		log.Fatalf("failed to connect db: %s", err)
	}
	defer conn.Close(context.Background())

}
