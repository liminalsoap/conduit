package middleware

import (
	"conduit-go/config"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
)

type MiddlewareManager struct {
	cfg    *config.Config
	log    logger.Interface
	userUC usecase.User
}

func NewMiddlewareManager(cfg *config.Config, log logger.Interface, userUC usecase.User) *MiddlewareManager {
	return &MiddlewareManager{cfg, log, userUC}
}
