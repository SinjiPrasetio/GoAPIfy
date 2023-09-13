package middleware

import (
	"GoAPIfy/config"
	"GoAPIfy/core"
	"GoAPIfy/model"
	"GoAPIfy/service/appService"
	"GoAPIfy/service/auth"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Authentication is a middleware that handles JWT token validation and user authentication.
// It extracts the JWT token from the "Authorization" header and validates it using the provided auth service.
// If the token is valid, it extracts the user ID from the token claims and fetches the corresponding user data
// from the model service. If the user data is valid, it sets it in the context for downstream handlers to access.
// If any errors occur during validation, it returns a 401 Unauthorized response with an error message.
func Authentication(authService auth.AuthService, s appService.AppService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			errorMessage := core.FormatError(errors.New("access denied : you're not authorized to call this api!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		tokenData, err := authService.ValidateToken(tokenString)
		if err != nil {
			errorMessage := core.FormatError(err)
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			c.Abort()
			return
		}

		// print raw claims
		fmt.Printf("Raw claims: %+v\n", tokenData.Claims)

		claimsPtr, ok := tokenData.Claims.(*jwt.MapClaims)
		if !ok {
			errorMessage := core.FormatError(errors.New("access denied : cannot extract token claims!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		// Dereference the pointer to get the actual claims
		claims := *claimsPtr

		sub, ok := claims["sub"]
		if !ok {
			errorMessage := core.FormatError(errors.New("access denied : user claim is missing!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		subFloat, ok := sub.(float64)
		if !ok {
			errorMessage := core.FormatError(errors.New("access denied : user claim is not a number!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		userID := uint(subFloat)

		var userModel model.User
		err = s.Model.Load(&userModel).Find(userID)
		if err != nil {
			errorMessage := core.FormatError(errors.New("access denied : user is unauthorized!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		if config.VerifyEmail() && userModel.VerifiedAt == nil {
			errorMessage := core.FormatError(errors.New("access denied : user email is not verified!"))
			core.SendResponse(c, http.StatusUnauthorized, errorMessage)
			return
		}

		c.Set("currentUser", userModel)
		c.Set("token", tokenString)
		c.Next()
	}
}
