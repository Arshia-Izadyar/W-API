package api

import (
	"fmt"
	"wapi/src/api/middleware"
	"wapi/src/api/routers"
	"wapi/src/api/validators"
	"wapi/src/config"
	"wapi/src/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.LimitByRequest())
	r.Use(middleware.Cors(cfg))

	RegisterRouts(r)
	RegisterSwagger(r, *cfg)
	RegisterValidators()

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))

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
		users := v1.Group("/users")
		routers.UserRouter(users, cfg)

		// test := v1.Group("/test")
		// routers.TestRouter(test)

		country := v1.Group("/country", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		routers.CountryRouter(country, cfg)

		cities := v1.Group("/city", middleware.Authentication(cfg))
		routers.CityRouter(cities, cfg)
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
