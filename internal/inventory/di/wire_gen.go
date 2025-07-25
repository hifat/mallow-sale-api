// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inventoryDi

import (
	"github.com/hifat/mallow-sale-api/internal/inventory/handler"
	"github.com/hifat/mallow-sale-api/internal/inventory/repository"
	"github.com/hifat/mallow-sale-api/internal/inventory/service"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func Init(cfg *config.Config, db *mongo.Database) *inventoryHandler.Handler {
	loggerLogger := logger.New()
	repository := inventoryRepository.NewMongo(db)
	usageUnitRepositoryRepository := usageUnitRepository.NewMongo(db)
	service := inventoryService.New(loggerLogger, repository, usageUnitRepositoryRepository)
	rest := inventoryHandler.NewRest(service)
	handler := inventoryHandler.New(rest)
	return handler
}
