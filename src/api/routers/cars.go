package routers

import (
	"wapi/src/api/handler"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func CarTypeRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarTypeHandler(cfg)
	r.GET("/get/:id", h.GetCarTypeById)
	r.POST("/create", h.CreateCarType)
	r.DELETE("/delete/:id", h.DeleteCarType)
	r.PUT("/update/:id", h.UpdateCarType)
	r.POST("/filter", h.GetByFilter)

}

func GearBoxRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewGearBoxHandler(cfg)
	r.GET("/get/:id", h.GetGearBoxById)
	r.POST("/create", h.CreateGearBox)
	r.DELETE("/delete/:id", h.DeleteGearBox)
	r.PUT("/update/:id", h.UpdateGearBox)
	r.POST("/filter", h.GetByFilter)

}

func CompanyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCompanyHandler(cfg)
	r.GET("/get/:id", h.GetCompanyById)
	r.POST("/create", h.CreateCompany)
	r.DELETE("/delete/:id", h.DeleteCompany)
	r.PUT("/update/:id", h.UpdateCompany)
	r.POST("/filter", h.GetByFilter)

}
