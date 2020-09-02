package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/command"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/user"
)

var Index = func(c *gin.Context) {
	text := c.PostForm("text")

	slackId := c.PostForm("user_id")
	user := user.GetOrCreateUser(slackId)

	response := command.ExecuteCommand(text, user)

	c.JSON(200, gin.H{"text": response})
}
