package entityModule

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updatedAt"`
}
