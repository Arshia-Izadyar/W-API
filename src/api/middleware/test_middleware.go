package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader("x-api-key")
		if key == "0910" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"result": "Key missing",
			})
		}
	}
}
