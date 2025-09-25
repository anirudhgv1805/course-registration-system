package middleware

import (
	"backend-crs/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  "Authorization Header is missing",
			})
			return
		}

		jwtToken := strings.TrimPrefix(authHeader, "Bearer ")
		userData, err := util.ValidateJwtToken(jwtToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  err.Error(),
			})
			return
		}

		ctx.Set("userId", userData["userId"])
		ctx.Set("role", userData["role"])
		fmt.Println("User data that is trying to access the resource ", userData)
		ctx.Next()
	}
}
