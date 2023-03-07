package user

import (
	"GoAPIfy/core"
	"GoAPIfy/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	modelService model.Model
}

func NewUserHandler(modelService model.Model) *userHandler {
	return &userHandler{modelService}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var input RegisterInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := core.FormatValidationErrors(err)

		errorMessage := gin.H{"errors": errors}

		core.GiveResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}
}
