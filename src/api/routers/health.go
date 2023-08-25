package routers

import (
	"wapi/src/api/handler"

	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	healthHandler := handler.NewHealthHandler()
	r.GET("/", healthHandler.HealthGet)

}
