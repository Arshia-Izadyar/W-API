package middleware

import (
	"net/http"
	"wapi/src/api/helper"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(gin.H{
				"Error": err.Error(),
			}, false, int(helper.LimiterError), err))
			return
		}
		ctx.Next()
	}
}
