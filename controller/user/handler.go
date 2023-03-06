package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ID uint
}

func User(c *gin.Context) {
	var data Data
	err := json.Unmarshal([]byte(`{"id":1}`), &data)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	c.JSON(http.StatusOK, data.ID)
}
