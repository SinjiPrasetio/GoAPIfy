package route

import (
	"GoAPIfy/controller"
	"GoAPIfy/model"

	"github.com/gin-gonic/gin"
)

func API(server *gin.Engine, modelService model.Model) {

	// Register all the handlers
	h := controller.RegisterHandler(modelService)

	// This is your API base path, you can rename it as you like.
	api := server.Group("/api/v1")

	// Define your routes here.
	api.POST("/test", h.UserHandler.Create)
	api.POST("/login", h.UserHandler.Login)
	// Add more routes as needed
}
