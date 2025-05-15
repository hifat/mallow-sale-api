package repoUtils

import (
	"errors"

	"github.com/hifat/cost-calculator-api/pkg/throw"
	"go.mongodb.org/mongo-driver/mongo"
)

// error
func MongoErr(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return throw.ErrRecordNotFound
	}

	return err
}
