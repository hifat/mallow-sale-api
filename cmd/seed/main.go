package main

import (
	"context"

	"github.com/hifat/cost-calculator-api/config"
	"github.com/hifat/cost-calculator-api/pkg/database"
	"github.com/hifat/cost-calculator-api/pkg/initial"
	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var logg core.Logger
var db *mongo.Database

func init() {
	cfg := config.LoadAppConfig("./", ".env")
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
