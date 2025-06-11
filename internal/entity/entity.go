package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        string     `fake:"{uuid}" json:"id" bson:"_id,omitempty"`
	CreatedAt *time.Time `fake:"{date}" json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `fake:"{date}" json:"updated_at" bson:"updated_at,omitempty"`
}

func (e *Base) SetDateTime() {
	now := time.Now()
	e.CreatedAt = &now
	e.UpdatedAt = &now
}

func (e *Base) SetUpdatedAt() {
	now := time.Now()
	e.UpdatedAt = &now
}

func (e *Base) SetID() {
	e.ID = primitive.NewObjectID().Hex()
}
