package bin

import (
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/challenge"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/user"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/userexercise"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*challenge.Challenge)(nil),
		(*user.User)(nil),
		(*userexercise.UserExercise)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
