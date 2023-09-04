package routers

import (
	"wapi/src/api/handler"
	"wapi/src/api/middleware"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {
	usersH := handler.NewUsersHandler(cfg)
	r.POST("/send-otp", middleware.OtpLimiter(cfg), usersH.SendOtp)
	r.POST("/login-by-username", usersH.LoginByUsername)
	r.POST("/register-login-by-phone", usersH.RegisterLoginByPhone)
	r.POST("/register-by-username", usersH.RegisterByUsername)

}
