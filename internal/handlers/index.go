package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/challenge"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/help"
)

type command struct {
	Text string `json:"text"`
}

var Index = func(c *gin.Context) {
	var command command
	c.BindJSON(&command)

	switch command.Text {
	case "challenge":
		c.JSON(200, gin.H{"text": challenge.GetHelp().Text})
		return
	case "help":
	default:
		c.JSON(200, gin.H{"text": help.GetHelp().Text})
	}
}
