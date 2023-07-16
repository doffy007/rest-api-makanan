package handler

import (
	"context"

	"github.com/doffy007/rest-api-makanan/config"
	"github.com/doffy007/rest-api-makanan/internal/repository"
	"github.com/doffy007/rest-api-makanan/internal/usecase"
)

type handler struct {
	ctx        context.Context
	conf       *config.Config
	usecase    usecase.Usecase
	repository repository.Repository
}

func NewHandler(ctx context.Context, conf *config.Config) Handler {
	return &handler{
		ctx:     ctx,
		conf:    conf,
		usecase: usecase.NewUsecase(ctx, conf),
	}
}
