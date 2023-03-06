package route

import (
	"GoAPI/controller/user"

	"github.com/gin-gonic/gin"
)

func API(server *gin.Engine) {
	// This is for your api base path, you can rename as you like.
	api := server.Group("/api/v1")

	// Define you route here
	api.POST("/test", user.User)

}
