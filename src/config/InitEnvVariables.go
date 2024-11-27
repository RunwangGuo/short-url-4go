package config

import (
	"os"
	"strconv"
	"time"
)

// HeaderTokenKey 定义安全码常量
const HeaderTokenKey = "X-Token"
const ValidToken = "53ROYinHId9qke"
const ShortIDLength = 5

type Config struct {
	Port          string
	Origin        string
	DBUsername    string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	Token         string
	CacheMaxCap   int
	CacheLiveTime time.Duration
	ApiSecret     string
	AccessLog     bool
}

func InitConfig() Config {
	return Config{
		Port:          getEnv("PORT", "80"),
		Origin:        getEnv("ORIGIN", "https://127.0.0.1"),
		DBUsername:    getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "root"),
		DBHost:        getEnv("DB_HOST", "127.0.0.1"),
		DBPort:        getEnv("DB_PORT", "3306"),
		DBName:        getEnv("DB_NAME", "short_url"),
		Token:         getEnv("TOKEN", "53ROYinHId9qke"),
		CacheMaxCap:   1000,
		CacheLiveTime: 60 * time.Second,
		ApiSecret:     getEnv("API_SECRET", "1FIsiEpxQo5l7H"),
		AccessLog:     getEnvAsBool("ACCESS_LOG", true),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// 获取 bool 类型的环境变量，如果不存在或解析失败则返回默认值
func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
