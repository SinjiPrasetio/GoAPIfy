// Package route provides the routing configuration for the API.
// It sets up the necessary middleware and handlers for the API endpoints.
package route

import (
	"GoAPIfy/controller"
	"GoAPIfy/middleware"
	"GoAPIfy/rate"
	"GoAPIfy/service/appService"
	"GoAPIfy/service/auth"

	"github.com/gin-gonic/gin"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

// API sets up the routes and middleware for the API.
// It takes a pointer to a gin.Engine instance and a model.Model instance.
// It registers the handlers and middleware required for the API, including authentication middleware,
// rate limiting using the Ulule limiter library, and the necessary routes for the API endpoints.
func API(server *gin.Engine, s appService.AppService) {
	authService := auth.NewJWTService(s)

	// Set up rate limiting
	rateLimit := rate.NewLimiter()
	server.Use(ginmiddleware.NewMiddleware(rateLimit))

	// Register all the handlers
	h := controller.RegisterHandler(s, authService)

	// This is your API base path, you can rename it as you like.
	api := server.Group("/api/v1")

	userGroup := api.Group("/user")

	// Define your routes here.
	userGroup.POST("/register", h.UserHandler.Register)
	userGroup.POST("/login", h.UserHandler.Login)

	// Define user group for routes that requires authentication
	userModGroup := userGroup.Group("/auth")

	// Use authentication middleware for routes that require authentication
	userModGroup.Use(middleware.Authentication(authService, s))

	// Add more routes as needed

}
