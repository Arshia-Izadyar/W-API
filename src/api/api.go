package api

import (
	"fmt"
	"wapi/src/api/routers"
	"wapi/src/api/validators"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	cfg := config.LoadCfg()
	api := r.Group("/api")
	vld, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		vld.RegisterValidation("phone", validators.IranPhoneNumberValidator, true)
		vld.RegisterValidation("password", validators.PassWordValidator, true)
	}
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)

		test := v1.Group("/test")
		routers.TestRouter(test)

	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))

}
