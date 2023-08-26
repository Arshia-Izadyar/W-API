package main

import (
	"wapi/src/api"
	"wapi/src/config"
	"wapi/src/data/cache"
	"wapi/src/data/db"
	"wapi/src/pkg/logging"
)

func main() {

	cfg := config.LoadCfg()

	logger := logging.NewLogger(cfg)

	// redis
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(err, logging.Redis, logging.Startup, "redis Failed:\t"+err.Error(), nil)
	}
	defer cache.CloseRedis()

	// postgres
	err = db.InitDB(*cfg)
	defer db.CloseDB()
	if err != nil {
		logger.Fatal(err, logging.Postgres, logging.Startup, "postgres Failed:"+err.Error(), nil)
	}

	// api server
	api.InitServer(cfg)
}
