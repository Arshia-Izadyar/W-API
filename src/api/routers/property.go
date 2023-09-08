package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func PropertyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyHandler(cfg)
	r.GET("/get/:id", h.GetProperty)
	r.POST("/create", h.CreateProperty)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteProperty)
	r.PUT("/update/:id", h.UpdateProperty)
}

func PropertyCategoryRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyCategoryHandler(cfg)
	r.GET("/get/:id", h.GetPropertyCategory)
	r.POST("/create", h.CreatePropertyCategory)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeletePropertyCategory)
	r.PUT("/update/:id", h.UpdatePropertyCategory)
}
