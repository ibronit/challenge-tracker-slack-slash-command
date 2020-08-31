package user

import (
	"time"

	"github.com/go-pg/pg/orm"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/exercise"
)

type User struct {
	ID        int
	SlackId   string
	Name      string
	CreatedAt time.Time `pg:"default:now()"`
	UpdatedAt time.Time `pg:"default:now()"`

	Exercises []exercise.Exercise `pg:"many2many:user_exercise"`
}

func (user *User) BeforeInsert(db orm.DB) error {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}

	return nil
}
