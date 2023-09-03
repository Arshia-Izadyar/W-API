package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {
	usersH := handler.NewUsersHandler(cfg)
	r.POST("send-otp", usersH.SendOtp)

}
