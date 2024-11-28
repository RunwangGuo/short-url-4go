package config

import (
	"github.com/caarlos0/env/v8"
	"time"
)

// HeaderTokenKey 定义安全码常量
const HeaderTokenKey = "X-Token"
const ValidToken = "53ROYinHId9qke"
const ShortIDLength = 5

type envConfig struct {
	Port       string `env:"PORT"`
	Origin     string `env:"ORIGIN"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBOptions  string `env:"DB_OPTIONS"`
	Token      string `env:"TOKEN"`
	//CacheMaxCapacity string `env:"CACHE_MAX_CAPACITY"`
	CacheMaxItems int           `env:"CACHE_MAX_ITEMS"`
	CacheLifetime time.Duration `env:"CACHE_LIFETIME"`
	ApiSecretKey  string        `env:"API_SECRET_KEY"`
	AccessLog     bool          `env:"ACCESS_LOG"`
}

var EnvVariables envConfig

func LoadEnvVariables() {
	if err := env.Parse(&EnvVariables); err != nil {
		panic("unable to load environment variables")
	}
}
