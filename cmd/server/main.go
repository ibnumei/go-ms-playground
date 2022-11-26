package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	route := gin.Default()

	route.Run(":" + os.Getenv("PORT"))
}