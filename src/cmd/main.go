package main

import (
	"log"
	"wapi/src/api"
	"wapi/src/config"
	"wapi/src/data/cache"
	"wapi/src/data/db"
)

func main() {

	// redis
	cfg := config.LoadCfg()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	// postgres
	err = db.InitDB(*cfg)
	defer db.CloseDB()
	if err != nil {
		log.Fatal(err)
	}

	// api server
	api.InitServer(cfg)
}
