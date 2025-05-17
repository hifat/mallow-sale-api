package router

import (
	"github.com/go-playground/validator/v10"
	core "github.com/hifat/goroger-core"
	"github.com/hifat/mallow-sale-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type router struct {
	route     core.IHttpRouter
	cfg       *config.Config
	db        *mongo.Database
	logger    *zap.Logger
	validator *validator.Validate
}

func New(route core.IHttpRouter, cfg *config.Config, db *mongo.Database, logger *zap.Logger, validator *validator.Validate) *router {
	return &router{
		route,
		cfg,
		db,
		logger,
		validator,
	}
}
