package routes

import (
	"github.com/kataras/iris/v12"
	"short-url-rw-github/src/controllers"
)

func RegisterRoutes(app *iris.Application, SearchController *controllers.SearchController, controller *controllers.ChangeExpiredController) {
	api := app.Party("/api")
	{
		api.Get("/search", SearchController.Search)      // 搜索链接
		api.Post("/status", linkController.ChangeStatus) // 状态更新接口
		api.Post("expired", controller.ChangeExpired)
	}
}
