package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/logger"
	"github.com/hifat/mallow-sale-api/config"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/initial"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var logg core.Logger

func dbConn(cfg config.Db) *mongo.Database {
	ctx := context.Background()

	dbClient := database.MongoConnect(ctx, &cfg)

	db := dbClient.Database(cfg.Name)

	logg = logger.New(initial.Logger)

	if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
		logg.Error(err)
	}

	return db
}

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Fatal("please provide all required parameters: ... <env_path> <service_name>")
	}

	cfg := config.LoadAppConfig(args[1])

	db := dbConn(cfg.Db)
	defer db.Client().Disconnect(context.Background())

	switch args[2] {
	case "usageUnit":
		if err := seedUsageUnit(db); err != nil {
			logg.Error(err)
		}
		slog.Info("seeded usage unit")
	}
}
