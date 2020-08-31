package user

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type User struct {
	ID        int
	SlackId   string
	Name      string
	CreatedAt time.Time `pg:"default:now()"`
	UpdatedAt time.Time

	Exercises []
}

func (user *User) BeforeInsert(db orm.DB) error {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	return nil
}
