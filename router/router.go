package router

import (
	"github.com/hifat/cost-calculator-api/config"
	core "github.com/hifat/goroger-core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type router struct {
	route  core.IHttpRouter
	cfg    *config.Config
	db     *mongo.Database
	logger *zap.Logger
}

func New(route core.IHttpRouter, cfg *config.Config, db *mongo.Database, logger *zap.Logger) *router {
	return &router{
		route,
		cfg,
		db,
		logger,
	}
}
