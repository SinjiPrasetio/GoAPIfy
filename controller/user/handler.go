package user

import (
	"GoAPIfy/model"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	modelService model.Model
}

func NewUserHandler(modelService model.Model) *userHandler {
	return &userHandler{modelService}
}

func (h *userHandler) CreateUser(c *gin.Context) {
}
