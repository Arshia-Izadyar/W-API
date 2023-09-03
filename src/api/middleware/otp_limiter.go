package middleware

import (
	"errors"
	"net/http"
	"time"
	"wapi/src/api/helper"
	"wapi/src/config"
	"wapi/src/pkg/limiter"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIpRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(ctx *gin.Context) {
		limiter := limiter.GetLimiter(ctx.ClientIP())
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, -1, errors.New("too many requests for otp")))
			ctx.Abort()
		} else {
			ctx.Next()
		}

	}
}
