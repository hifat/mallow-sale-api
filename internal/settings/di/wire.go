//go:build wireinject
// +build wireinject

package settingDi

import (
	"github.com/google/wire"
	settingsHandler "github.com/hifat/mallow-sale-api/internal/settings/handler"
	settingRepository "github.com/hifat/mallow-sale-api/internal/settings/repository"
	settingService "github.com/hifat/mallow-sale-api/internal/settings/service"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeSettingsRest(db *mongo.Database) *settingsHandler.Rest {
	wire.Build(
		/* ------------------------------- Repository ------------------------------- */
		settingRepository.NewMongo,

		/* --------------------------------- Service -------------------------------- */
		logger.New,
		settingService.New,

		/* --------------------------------- Handler -------------------------------- */
		settingsHandler.NewRest,
	)
	return nil
}
