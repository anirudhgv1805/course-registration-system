package middleware

import (
	"backend-crs/util"
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
		registerNo, err := util.ValidateJwtToken(jwtToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  err.Error(),
			})
			return
		}

		ctx.Set("registerNo", registerNo)
		ctx.Next()
	}
}
