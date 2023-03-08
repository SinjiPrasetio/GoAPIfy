package route

import (
	"GoAPIfy/controller"
	"GoAPIfy/model"
	"GoAPIfy/service/auth"

	"github.com/gin-gonic/gin"
)

func API(server *gin.Engine, modelService model.Model) {
	authService := auth.NewService()

	// Register all the handlers
	h := controller.RegisterHandler(modelService, authService)

	// This is your API base path, you can rename it as you like.
	api := server.Group("/api/v1")

	// Define your routes here.
	api.POST("/user/create", h.UserHandler.Create)
	api.POST("/user/login", h.UserHandler.Login)
	// Add more routes as needed
}
