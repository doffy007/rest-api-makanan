package repository

import (
	"context"

	"github.com/doffy007/rest-api-makanan/config"
	"github.com/doffy007/rest-api-makanan/database"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	ctx  context.Context
	conf *config.Config
	db   *sqlx.DB
}

func NewRepository(ctx context.Context, conf *config.Config) Repository {
	return &repository{
		ctx:  ctx,
		conf: conf,
		db:   database.Mysql(),
	}
}
