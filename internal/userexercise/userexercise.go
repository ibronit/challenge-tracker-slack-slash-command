package userexercise

import (
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/exercise"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/user"
)

type UserExercise struct {
	tableName struct{} `pg:"alias:bg"`

	UserId     int `pg:",pk"`
	User       *user.User
	ExerciseId int `pg:",pk"`
	Exercise   *exercise.Exercise

	amount int
}
