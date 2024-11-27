package controllers

import (
	"github.com/kataras/iris/v12"
	"log"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
)

type GenerateController struct {
	interfaces.IGenerateService
	EnvVariables *config.Config
}

func (g *GenerateController) GenerateController(ctx iris.Context) {
	// 校验 Token
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" || headerToken != g.EnvVariables.Token {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("请提供正确的安全码")
		return
	}

	// 解析请求体
	var params models.GenerateReq
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("参数解析错误")
		return
	}

	// 调用Service处理逻辑
	results, err := g.IGenerateService.Generate(params.URLs, params.ExpiredTs)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("生成短链接失败")
		return
	}

	// 返回结果
	ctx.ContentType("application/json")
	if err := ctx.JSON(results); err != nil {
		log.Printf("返回JSON失败：%v\n", err)
	}

}
