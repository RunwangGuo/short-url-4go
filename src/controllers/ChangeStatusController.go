package controllers

import (
	"github.com/kataras/iris/v12"
	"log"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
)

type ChangeStatus struct {
	interfaces.ILinkService
	EnvVariables *config.Config
}

func (c *ChangeStatus) ChangeStatusController(ctx iris.Context) {
	// 验证Token
	headerToken := ctx.GetHeader("Authorization")
	if headerToken != c.EnvVariables.Token {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("请求参数错误")
		return
	}

	// 解析请求体
	var req models.ChangeStatusReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("请求参数错误")
		return
	}

	// 调用service更新状态
	err := c.UpdateStatus(req.Targets, req.Status)
	if err != nil {
		log.Printf("ChangeStatus error: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.WriteString("状态更新失败")
		return
	}

	// 返回成功响应
	err = ctx.JSON(iris.Map{"message": "状态更新成功"})
	if err != nil {
		return
	}
}
