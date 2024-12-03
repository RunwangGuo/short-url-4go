package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"go.uber.org/zap"
	"time"
)

// HeaderTokenKey 定义安全码常量
const HeaderTokenKey = "X-Token"
const ShortIDLength = 5

type envConfig struct {
	Port       string `env:"PORT" default:"8080"`
	Origin     string `env:"ORIGIN" default:"http://localhost"`
	DBHost     string `env:"DB_HOST" default:"127.0.0.1"`
	DBPort     string `env:"DB_PORT" default:"3306"`
	DBUsername string `env:"DB_USERNAME" default:"root"`
	DBPassword string `env:"DB_PASSWORD" default:"root"`
	DBName     string `env:"DB_NAME" default:"short_url"`
	DBOptions  string `env:"DB_OPTIONS"`
	//CacheMaxCapacity string `env:"CACHE_MAX_CAPACITY"`
	CacheMaxItems int           `env:"CACHE_MAX_ITEMS"`
	CacheLifetime time.Duration `env:"CACHE_LIFETIME"`
	AccessLog     bool          `env:"ACCESS_LOG"`
	Token         string        `env:"TOKEN"  default:"53ROYinHId9qke"`
	ApiSecret     string        `env:"API_SECRET" default:"1FIsiEpxQo5l7H"`
}

var EnvVariables envConfig

func LoadEnvVariables() {
	if err := env.Parse(&EnvVariables); err != nil {
		panic("unable to load environment variables")
	}
	// 打印加载结果以调试
	//fmt.Printf("Loaded Config: %+v\n", EnvVariables)
	// 使用 fmt.Sprintf 来格式化打印
	envInfo := fmt.Sprintf("Loaded Config: Port=%s, Origin=%s, DBHost=%s, DBPort=%s, DBUsername=%s, DBName=%s, CacheMaxItems=%d, CacheLifetime=%s, Token=%s, ApiSecret=%s, AccessLog=%v",
		EnvVariables.Port,
		EnvVariables.Origin,
		EnvVariables.DBHost,
		EnvVariables.DBPort,
		EnvVariables.DBUsername,
		EnvVariables.DBName,
		EnvVariables.CacheMaxItems,
		EnvVariables.CacheLifetime,
		EnvVariables.Token,
		EnvVariables.ApiSecret,
		EnvVariables.AccessLog,
	)

	// 用 ZapLogger 输出环境变量信息
	ZapLogger.Info("EnvVariables初始化成功", zap.String("config", envInfo))
}
