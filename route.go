package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"io"
	"os"
	"short-url-4go/src/interfaces"
	"short-url-4go/src/middleware"
	"sync"
)

var (
	irisRouter *router
	routerOnce sync.Once
)

type router struct{}

type IRouter interface {
	InitRouter(dbClient interfaces.IDataAccessLayer, cache interfaces.ICacheLayer) *iris.Application
}

func (router *router) InitRouter(dbClient interfaces.IDataAccessLayer, cache interfaces.ICacheLayer) *iris.Application {
	app := iris.New()

	// Our custom CORS middleware.
	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if ctx.Method() == iris.MethodOptions {
			ctx.Header("Access-Control-Methods",
				"POST, PUT, PATCH, DELETE")

			ctx.Header("Access-Control-Allow-Headers",
				"Access-Control-Allow-Origin,Content-Type,Authorization,api-secret,token")

			ctx.Header("Access-Control-Max-Age",
				"86400")

			ctx.StatusCode(iris.StatusNoContent)
			return
		}

		ctx.Next()
	}

	app.UseRouter(crs)

	LinkController := ServiceContainer().InjectLinkController(dbClient, cache)

	//app.Get("/healthcheck", healthCheckController.CheckServerHealthCheck)
	//app.Post("/v1/shorten", middleware.CheckJWT(), shortifyWriterController.WriterController)
	app.Post("/api/redirect", LinkController.Redirect)
	app.Get("/api/search", middleware.CheckApiSecret(), LinkController.Search)
	app.Post("/api/generate", middleware.CheckApiSecret(), middleware.CheckToken(), LinkController.Generate)
	app.Post("/api/change_status", LinkController.ChangeStatus)
	app.Post("/api/change_expired", LinkController.ChangeExpired)
	app.Post("/api/remark", LinkController.Remark)

	return app
}

func Router() IRouter {
	if irisRouter == nil {
		routerOnce.Do(func() {
			irisRouter = &router{}
		})
	}
	return irisRouter
}

// This helps to log the request and its metadata
func makeAccessLog() *accesslog.AccessLog {
	ac := accesslog.New(io.MultiWriter(os.Stdout))
	ac.Delim = ' '
	ac.ResponseBody = false
	ac.RequestBody = false
	ac.BytesReceived = true
	ac.BytesSent = true

	return ac
}
