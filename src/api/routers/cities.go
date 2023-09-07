package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func CityRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCityHandler(cfg)
	r.GET("/get/:id", h.GetCityById)
	r.POST("/create", h.CreateCity) // crud
	r.PUT("/update/:id", h.UpdateCity)
	r.DELETE("/delete/:id", h.DeleteCity)
	r.POST("/filter", h.GetByFilter)

}
