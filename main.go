package main

import (
	"os"

	"github.com/ibronit/challenge-tracker-slack-slash-command/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/handlers"
)

func main() {
	router := gin.Default()

	db.GetDB()

	router.POST("/", handlers.Index)

	router.Run(":" + os.Getenv("PORT"))
}
