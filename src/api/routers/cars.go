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

func CarModelRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelHandler(cfg)
	r.GET("/get/:id", h.GetCarModelById)
	r.POST("/create", h.CreateCarModel)
	r.DELETE("/delete/:id", h.DeleteCarModel)
	r.PUT("/update/:id", h.UpdateCarModel)
	r.POST("/filter", h.GetByFilter)

}

func CarModelColorRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelColorHandler(cfg)
	r.GET("/get/:id", h.GetCarModelColorById)
	r.GET("/get-c/:id", h.GetCarModelColorById)
	r.POST("/create", h.CreateCarModelColor)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteCarModelColor)
	r.PUT("/update/:id", h.UpdateCarModelColor)
}

func PersianYearRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPersianYearHandler(cfg)
	r.GET("/get/:id", h.GetPersianYearById)
	r.POST("/create", h.CreatePersianYear)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeletePersianYear)
	r.PUT("/update/:id", h.UpdatePersianYear)
}

func CarModelYearRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelYearHandler(cfg)
	r.GET("/get/:id", h.GetCarModelYearById)
	r.POST("/create", h.CreateCarModelYear)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteCarModelYear)
	r.PUT("/update/:id", h.UpdateCarModelYear)
}

func CarModelPriceRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelPriceHandler(cfg)
	r.GET("/get/:id", h.GetCarModelPriceById)
	r.POST("/create", h.CreateCarModelPrice)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteCarModelPrice)
	r.PUT("/update/:id", h.UpdateCarModelPrice)
}

func CarModelFileRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelFileHandler(cfg)
	r.GET("/get/:id", h.GetCarModelFileById)
	r.POST("/create", h.CreateCarModelFile)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteCarModelFile)
	r.PUT("/update/:id", h.UpdateCarModelFile)
}

func CarModelPropertyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelPropertyHandler(cfg)
	r.GET("/get/:id", h.GetCarModelPropertyById)
	r.POST("/create", h.CreateCarModelProperty)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteCarModelProperty)
	r.PUT("/update/:id", h.UpdateCarModelProperty)
}

func CarModelCommentRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelCommentHandler(cfg)
	r.GET("/get/:id", h.GetCarModelCommentById)
	r.POST("/create", h.CreateCarModelComment)
	r.POST("/filter", h.GetByFilter)
	r.DELETE("/delete/:id", h.DeleteCarModelComment)
	r.PUT("/update/:id", h.UpdateCarModelComment)
}
