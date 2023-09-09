package api

import (
	"fmt"
	"wapi/src/api/middleware"
	"wapi/src/api/routers"
	"wapi/src/api/validators"
	"wapi/src/config"
	"wapi/src/docs"
	"wapi/src/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.LoadCfg())

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.LimitByRequest())
	r.Use(middleware.Cors(cfg))

	RegisterRouts(r)
	RegisterSwagger(r, *cfg)
	RegisterValidators()

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		logger.Error(err, logging.General, logging.Startup, "server run failed", nil)
	}

}
func RegisterValidators() {
	vld, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		vld.RegisterValidation("phone", validators.IranPhoneNumberValidator, true)
		vld.RegisterValidation("password", validators.PassWordValidator, true)
	}
}

func RegisterRouts(r *gin.Engine) {
	cfg := config.LoadCfg()
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// test
		// test := v1.Group("/test")
		// routers.TestRouter(test)

		// users
		users := v1.Group("/users")
		routers.UserRouter(users, cfg)

		// countries
		country := v1.Group("/country")
		routers.CountryRouter(country, cfg)

		// cities
		cities := v1.Group("/city", middleware.Authentication(cfg))
		routers.CityRouter(cities, cfg)

		// files
		files := v1.Group("/files", middleware.Authentication(cfg))
		routers.FileRouter(files, cfg)

		// property
		property := v1.Group("/property", middleware.Authentication(cfg))
		routers.PropertyRouter(property, cfg)

		// car
		companies := v1.Group("/company", middleware.Authentication(cfg))
		routers.CompanyRouter(companies, cfg)

		carType := v1.Group("/car-type", middleware.Authentication(cfg))
		routers.CarTypeRouter(carType, cfg)

		gearbox := v1.Group("/gearbox", middleware.Authentication(cfg))
		routers.GearBoxRouter(gearbox, cfg)

		propertyCategory := v1.Group("/property-category", middleware.Authentication(cfg))
		routers.PropertyCategoryRouter(propertyCategory, cfg)

		carModel := v1.Group("/car-model", middleware.Authentication(cfg))
		routers.CarModelRouter(carModel, cfg)

		carModelColor := v1.Group("/car-color", middleware.Authentication(cfg))
		routers.CarModelColorRouter(carModelColor, cfg)

		Color := v1.Group("/color", middleware.Authentication(cfg))
		routers.ColorRouter(Color, cfg)

		Year := v1.Group("/year", middleware.Authentication(cfg))
		routers.PersianYearRouter(Year, cfg)

		carYear := v1.Group("/car-year", middleware.Authentication(cfg))
		routers.CarModelYearRouter(carYear, cfg)

		carPrice := v1.Group("/car-price", middleware.Authentication(cfg))
		routers.CarModelPriceRouter(carPrice, cfg)
	}
}

func RegisterSwagger(r *gin.Engine, cfg config.Config) {

	docs.SwaggerInfo.Title = "golang first api"
	docs.SwaggerInfo.Description = "golang first api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
