package route

import (
	"GoAPIfy/controller"
	"GoAPIfy/middleware"
	"GoAPIfy/model"
	"GoAPIfy/service/auth"

	"github.com/gin-gonic/gin"
)

func API(server *gin.Engine, modelService model.Model) {
	authService := auth.NewJWTService()

	// Register all the handlers
	h := controller.RegisterHandler(modelService, authService)

	// This is your API base path, you can rename it as you like.
	api := server.Group("/api/v1")

	userGroup := api.Group("/user")

	// Define your routes here.
	userGroup.POST("/register", h.UserHandler.Register)
	userGroup.POST("/login", h.UserHandler.Login)

	userModGroup := userGroup.Group("/edit")

	userModGroup.Use(middleware.Authentication(authService, modelService))
	// Add more routes as needed
}
