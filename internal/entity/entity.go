package entity

import (
	"time"
)

type Base struct {
	ID        string    `bson:"_id,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (e *Base) SetDateTime() {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
}
