package main

import (
	"context"
	"fmt"

	"github.com/hifat/mallow-sale-api/internal/entity"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func seedUsageUnit(db *mongo.Database) error {
	ctx := context.Background()

	docs := []usageUnit.UsageUnit{
		{
			Code: "ML",
			Name: "ml",
			Base: entity.Base{
				CreatedAt: utils.TimeNow(),
				UpdatedAt: utils.TimeNow(),
			},
		},
		{
			Code: "G",
			Name: "g",
			Base: entity.Base{
				CreatedAt: utils.TimeNow(),
				UpdatedAt: utils.TimeNow(),
			},
		},
		{
			Code: "KG",
			Name: "kg",
			Base: entity.Base{
				CreatedAt: utils.TimeNow(),
				UpdatedAt: utils.TimeNow(),
			},
		},
	}

	_usageUnit := usageUnit.UsageUnit{}
	newDocs := make([]interface{}, 0, len(docs))
	for _, doc := range docs {
		newDocs = append(newDocs, doc)
	}

	result, err := db.Collection(_usageUnit.Doc()).
		InsertMany(ctx, newDocs)
	if err != nil {
		return err
	}

	logg.Info(fmt.Sprintf("seeded UsageUnit: %v", result.InsertedIDs))

	return nil
}
