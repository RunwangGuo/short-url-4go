package main

import (
	"short-url-4go/src/config"
	_ "short-url-4go/src/infrastrctures"
)

func main() {
	//initialize logger
	config.InitializeLogger()

	//initialize env variables
	config.LoadEnvVariables()

	//initialize DB
	config.MySQL().InitMySQLConnection()
	config.MySQL().InitTables()

	//initialize gcache
	config.Gcache().InitGCache()

	//initialize api routes
	//app := iris.New()
	//app := Router().InitRouter(config.DynamoDB().(*config.DBClientHandler).DBClient, config.Redis().(*config.RedisHandler).RedisClient)
	app := Router().InitRouter(config.MySQL().(*config.MySQLHandler).DBClient, config.Gcache().(*config.GCacheHandler).Cache)

	// 启动服务
	err := app.Listen(":" + config.EnvVariables.Port)
	if err != nil {
		panic("unable to start server")
	}
}
