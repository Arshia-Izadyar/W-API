package api

import (
	"fmt"
	"wapi/src/api/middleware"
	"wapi/src/api/routers"
	"wapi/src/api/validators"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.LimitByRequest())
	r.Use(middleware.Cors(cfg))

	RegisterRouts(r)
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
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)

		test := v1.Group("/test")
		routers.TestRouter(test)

	}
}
