package command

import (
	"fmt"

	"github.com/alexflint/go-restructure"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/challenge"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/user"
)

type Command struct {
	_      struct{} `^`
	Entity string   `\w+`
	_      struct{} `\s`
	Action string   `[^\s]\w+`
	_      struct{} `$?`
}

type Parameter struct {
	_              struct{}        `[-]{2}`
	Parameter      string          `\w+`
	ParameterValue *ParameterValue `?`
}

type ParameterValue struct {
	_     struct{} `[=]`
	Value string   `\w+`
}

func ExecuteCommand(commandText string, user *user.User) string {
	var command Command
	restructure.Find(&command, commandText)

	var parameters []Parameter
	restructure.MustCompile(Parameter{}, restructure.Options{}).FindAll(&parameters, commandText, -1)

	switch command.Entity {
	case "challenge":
		for _, value := range parameters {
			fmt.Println(value.Parameter)

			if value.ParameterValue != nil {
				fmt.Println(value.ParameterValue.Value)
			}
		}
		return challenge.GetHelp().Text
	default:
		return GetHelp().Text
	}
}
