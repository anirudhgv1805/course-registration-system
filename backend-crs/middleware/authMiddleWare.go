package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRoles(acceptedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requesterRole, isPresent := ctx.Get("role")
		if !isPresent {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  "Access forbidden",
			})
			return
		}
		for _, role := range acceptedRoles {
			if role == requesterRole {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  "Access forbidden",
		})
	}
}
