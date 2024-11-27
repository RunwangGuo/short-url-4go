package controllers

import (
	"github.com/kataras/iris/v12"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
)

type RemarkController struct {
	interfaces.ILinkService
	EnvVariables *config.Config
}

// Remark 修改备注的控制器
func (r *RemarkController) Remark(ctx iris.Context) {
	// 校验 Token
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" || headerToken != r.EnvVariables.Token {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("请提供正确的安全码")
		return
	}

	// 解析请求体
	var params models.RemarkReq
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("参数解析错误")
		return
	}

	// 调用Service处理逻辑
	//results, err := r.IGenerateService.Generate(params.URLs, params.ExpiredTs)
	err := r.ILinkService.UpdateRemark(params.Targets, params.Remark)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("更新备注失败")
		return
	}

	// 返回成功响应
	ctx.ContentType("application/json")
	ctx.WriteString("{}")
}
