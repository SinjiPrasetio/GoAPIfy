package middleware

import (
	"Laravel/core"
	"Laravel/model"
	"Laravel/service/auth"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication(authService auth.AuthService, modelService model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			errorMessage := core.FormatError(errors.New("access denied : you're not authorized to call this api!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		// Split Bearer dan Token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			errorMessage := core.FormatError(errors.New("access denied : fail to validate token!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			errorMessage := core.FormatError(errors.New("access denied : token is not valid!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		userID := uint(claim["id"].(float64))

		var userModel model.User
		result, err := modelService.FindByID(userID, userModel)
		if err != nil {
			errorMessage := core.FormatError(errors.New("access denied : user is unauthorized!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}
		userData, ok := result.(model.User)
		if !ok {
			errorMessage := core.FormatError(errors.New("access denied : user data corrupted!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		c.Set("currentUser", userData)
		c.Next()
	}
}
