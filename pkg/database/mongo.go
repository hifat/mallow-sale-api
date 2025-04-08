package database

import (
	"context"
	"fmt"
	"time"

	"github.com/hifat/cost-calculator-api/config"
	"github.com/hifat/cost-calculator-api/pkg/initial"
	"github.com/hifat/goroger-core/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoConnect(pctx context.Context, cfg *config.Db) *mongo.Client {
	if initial.Logger == nil {
		panic("logger is not initial")
	}

	log := logger.New(initial.Logger)

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func MustStrToObjectID(hex string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(hex)

	return objectID
}
