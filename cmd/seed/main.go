package main

import (
	"context"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/logger"
	"github.com/hifat/mallow-sale-api/config"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/initial"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var logg core.Logger
var db *mongo.Database

func init() {
	cfg := config.LoadAppConfig("./.env")
	ctx := context.Background()

	dbClient := database.MongoConnect(ctx, &cfg.Db)
	// defer dbClient.Disconnect(ctx)

	db = dbClient.Database(cfg.Db.Name)

	logg = logger.New(initial.Logger)

	if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
		logg.Error(err)
	}
}

func main() {
	defer db.Client().Disconnect(context.Background())

	if err := seedUsageUnit(); err != nil {
		logg.Error(err)
	}
}
