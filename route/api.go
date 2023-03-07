package route

import (
	"GoAPIfy/controller/user"
	"GoAPIfy/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func API(server *gin.Engine, db *gorm.DB) {
	// Loading modelService
	modelService := model.NewModel(db)

	//Define your handler here...
	userHandler := user.NewUserHandler(modelService)

	// This is your API base path, you can rename it as you like.
	api := server.Group("/api/v1")

	// Define your routes here.
	api.POST("/test", userHandler.CreateUser)
}
