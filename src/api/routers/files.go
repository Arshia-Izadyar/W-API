package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func FileRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewFileHandler(cfg)

	r.GET("/get/:id", h.GetFile)
	r.POST("/create", h.CreateFile)
	r.PUT("/update/:id", h.UpdateFile)
	r.DELETE("/delete/:id", h.DeleteFile)
}
