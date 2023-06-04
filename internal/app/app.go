package app

import (
	"conduit-go/config"
	"conduit-go/internal/delivery/http"
	"conduit-go/internal/repository"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"conduit-go/pkg/postgres"
	"context"
	"github.com/gin-gonic/gin"
)

func initUseCases(pg *postgres.Postgres) usecase.UseCases {
	var useCases usecase.UseCases
	tagUseCase := usecase.NewTagUseCase(
		repository.NewTagRepo(pg),
	)
	userUseCase := usecase.NewUserUseCase(
		repository.NewUserRepo(pg),
	)
	useCases.Tag = tagUseCase
	useCases.User = userUseCase
	return useCases
}

func Run(cfg *config.Config) {
	log := logger.NewLogger(cfg.Logger.Level)

	pg, err := postgres.NewDb(cfg.Postgresql.Url)
	if err != nil {
		log.Fatalf("failed to connect db: %s", err)
	}
	defer pg.Conn.Close(context.Background())

	useCases := initUseCases(pg)

	handler := gin.Default()
	http.NewRouter(handler, log, useCases)
	log.Fatalf("router error: %s", handler.Run(cfg.Http.Port))
}
