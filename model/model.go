package model

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-pg/pg/orm"
)

// Base contains common fields for all tables
type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListQuery holds company/location data used for list db queries
type ListQuery struct {
	Query string
	ID    string
}

// ListQuery holds company/location data used for list db queries
type FilterQuery struct {
	Query  string
	Params []interface{}
}

// BeforeInsert hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeInsert(c context.Context, db orm.DB) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now

	return nil
}

// BeforeUpdate hooks into update operations, setting updatedAt to current time
func (b *Base) BeforeUpdate(c context.Context, db orm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Base) ToString() string {
	str, _ := json.Marshal(b)
	return string(str)
}
