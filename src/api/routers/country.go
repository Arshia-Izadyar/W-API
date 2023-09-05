package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func CountryRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCountryHandler(cfg)
	r.GET("/get-by-id/:id", h.GetCountryById)
	r.POST("/create", h.CreateCountry)
	r.DELETE("/delete/:id", h.DeleteCountry)
	r.PUT("/update/:id", h.UpdateCountry)
}