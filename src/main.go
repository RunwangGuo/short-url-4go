package main

import (
	"github.com/kataras/iris/v12"
	"log"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/controllers"
)

func main() {
	// 初始化配置、数据库和缓存
	cfg := config.InitConfig()
	config.InitDB(cfg)
	config.InitCache(cfg)

	app := iris.New()

	generateController := controllers.GenerateController{IGenerateService, EnvVariables}

	// 路由设置
	app.Get("/{short_id:string}", controllers.Redirect)
	app.Post("/api/generate", generateController.GenerateController(ctx))
	app.Post("/api/status", controllers.ChangeStatus)
	app.Post("/api/expired", controllers.ChangeExpired)

	// 启动服务
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

/*	// 初始化存储库
	linkRepo := repositories.NewLinkRepository(db)
	cacheRepo := repositories.NewCacheRepository(1000) // 使用 gcache，容量为 1000

	// 初始化服务
	linkService := services.NewLinkService(linkRepo, cacheRepo)

	// 初始化控制器
	controller := &controllers.ChangeExpiredController{
		Service: linkService,
		Token:   "your-secret-token",
	}

	// 创建 Iris 应用
	app := iris.New()

	// 注册路由
	routes.RegisterRoutes(app, controller)*/
