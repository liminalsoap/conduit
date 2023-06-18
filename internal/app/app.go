package app

import (
	"conduit-go/config"
	"conduit-go/internal/delivery/http"
	"conduit-go/internal/middleware"
	"conduit-go/internal/repository"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"conduit-go/pkg/postgres"
	"context"
	"github.com/gin-gonic/gin"
)

func initUseCases(pg *postgres.Postgres) usecase.UseCases {
	var useCases usecase.UseCases
	userRepo := repository.NewUserRepo(pg)
	tagUseCase := usecase.NewTagUseCase(
		repository.NewTagRepo(pg),
	)
	userUseCase := usecase.NewUserUseCase(
		userRepo,
	)
	followingUceCase := usecase.NewFollowUseCase(
		repository.NewFollowingRepo(pg),
		userRepo,
	)
	articleTagUseCase := usecase.NewArticleTagUseCase(
		repository.NewArticleTagRepo(pg),
	)
	articleUseCase := usecase.NewArticleUseCase(
		repository.NewArticleRepo(pg),
		tagUseCase,
		articleTagUseCase,
	)
	likeUseCase := usecase.NewLikeUseCase(
		repository.NewLikeRepo(pg),
	)
	commentUseCase := usecase.NewCommentUseCase(
		repository.NewCommentRepo(pg),
		articleUseCase,
	)
	useCases.Tag = tagUseCase
	useCases.User = userUseCase
	useCases.Following = followingUceCase
	useCases.Article = articleUseCase
	useCases.Like = likeUseCase
	useCases.Comment = commentUseCase
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
	mw := middleware.NewMiddlewareManager(cfg, log, useCases.User)
	http.NewRouter(handler, log, useCases, cfg, mw)
	log.Fatalf("router error: %s", handler.Run(cfg.Http.Port))
}
