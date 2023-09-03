package config

import (
	"fmt"
	"log"
	time "time"

	viper "github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Cors     CorsConfig
	Logger   LoggerConfig
	PassWord PasswordConfig
	Otp      OtpConfig
}
type OtpConfig struct {
	Digits     int
	ExpireTime time.Duration
	Limiter    time.Duration
}

type CorsConfig struct {
	AllowOrigins string
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SslMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               int
	Password           string
	Db                 int
	PoolSize           int
	MinIdleConnections int
	PoolTimeout        time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "../config/config-docker.yml"
	} else if env == "development" {
		return "../config/config-development.yml"
	} else {
		return "../config/config-development.yml"
	}
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("cant open file %s", err.Error())
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("Viper: ConfigFileNotFoundErrorFile not found: %s.%s", fileName, fileType)
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadCfg() *Config {
	cfgPath := getConfigPath("APP_ENV")
	v, err := LoadConfig(cfgPath, "yaml")
	if err != nil {
		log.Fatal("Error in load config", err)
	}
	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Error in parse cfg %s", err)
	}

	return cfg
}
