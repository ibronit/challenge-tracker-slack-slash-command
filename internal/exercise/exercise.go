package exercise

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type Exercise struct {
	ID        int
	name      string
	CreatedAt time.Time `pg:"default:now()"`
	UpdatedAt time.Time
}

func (exercise *Exercise) BeforeInsert(db orm.DB) error {
	if exercise.CreatedAt.IsZero() {
		exercise.CreatedAt = time.Now()
	}
	return nil
}
