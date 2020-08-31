package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/challenge"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/help"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/user"
	"github.com/ibronit/challenge-tracker-slack-slash-command/pkg/db"
)

var Index = func(c *gin.Context) {
	text := c.PostForm("text")

	switch text {
	case "challenge":

		slackId := c.PostForm("user_id")

		user1 := &user.User{
			Name:    "admin",
			SlackId: slackId,
		}

		db.GetDB().Model(user1).Insert()

		c.JSON(200, gin.H{"text": challenge.GetHelp().Text})
		return
	default:
		c.JSON(200, gin.H{"text": help.GetHelp().Text})

	}
}
