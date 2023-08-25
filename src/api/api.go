package api

import (
	"fmt"
	"wapi/src/api/routers"
	"wapi/src/config"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	cfg := config.LoadCfg()
	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)

	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))

}
