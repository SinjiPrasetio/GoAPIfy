package route

import (
	"GoAPI/controller/user"

	"github.com/gin-gonic/gin"
)

func API(server *gin.Engine) {
	// This is your API base path, you can rename it as you like.
	api := server.Group("/api/v1")

	// Define your routes here.
	api.POST("/test", user.CreateUser)
}
