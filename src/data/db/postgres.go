package db

import (
	"fmt"
	"time"
	"wapi/src/config"
	"wapi/src/pkg/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var cfg = config.LoadCfg()
var logger = logging.NewLogger(cfg)

func InitDB(cfg config.Config) (err error) {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
	)
	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := dbClient.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)
	logger.Info(logging.Postgres, logging.Startup, "postgres started", nil)
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	db, _ := dbClient.DB()
	db.Close()
}
