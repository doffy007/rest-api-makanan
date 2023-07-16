package router

import (
	"context"

	"github.com/doffy007/rest-api-makanan/config"
	"github.com/doffy007/rest-api-makanan/internal/handler"
	"github.com/gorilla/mux"
)

type router struct {
	ctx     context.Context
	config  *config.Config
	route   *mux.Router
	handler handler.Handler
}

func Register(ctx context.Context, conf *config.Config, route *mux.Router) Router {
	return &router{
		ctx:     ctx,
		config:  conf,
		route:   route,
		handler: handler.NewHandler(ctx, conf),
	}
}

func (r router) All() {
	r.BaseRouter()
}
