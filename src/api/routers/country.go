package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func CountryRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCountryHandler(cfg)
	r.GET("/get/:id", h.GetCountryById)
	r.GET("/get-c/:id", h.GetCitiesById)
	r.POST("/create", h.CreateCountry)
	r.POST("/filter", h.GetCountryByFilter)
	r.DELETE("/delete/:id", h.DeleteCountry)
	r.PUT("/update/:id", h.UpdateCountry)
}
