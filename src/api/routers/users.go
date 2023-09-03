package routers

import (
	"wapi/src/api/handler"
	"wapi/src/api/middleware"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {
	usersH := handler.NewUsersHandler(cfg)
	r.POST("send-otp", middleware.OtpLimiter(cfg), usersH.SendOtp)

}
