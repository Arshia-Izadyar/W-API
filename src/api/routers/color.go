package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func ColorRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewColorHandler(cfg)
	r.GET("/get/:id", h.GetColorById)
	r.GET("/get-c/:id", h.GetColorById)
	r.POST("/create", h.CreateColor)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteColor)
	r.PUT("/update/:id", h.UpdateColor)
}
