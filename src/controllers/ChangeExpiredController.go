package controllers

import (
	"github.com/kataras/iris/v12"
	"log"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
	"time"
)

type ChangeExpiredController struct {
	interfaces.ILinkService
	EnvVariables *config.Config
}

// ChangeExpired 修改过期时间的控制器
func (c *ChangeExpiredController) ChangeExpired(ctx iris.Context) {

	// 校验Token
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" || headerToken != c.EnvVariables.Token {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("请提供正确的安全码")
		return
	}

	// 解析请求体
	var params models.ChangeExpiredReq
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("参数解析错误")
		return
	}

	// 校验过期时间是否合理
	if params.Expired < time.Now().Unix() {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("请提供不小于当前日期的过期时间")
		return
	}

	/*	// 校验时间戳是否合理
		if !utils.IsReasonableTimestamp(params.Expired) {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.WriteString("请提供不小于当前日期的过期时间")
			return
		}*/

	// 调用Service处理业务逻辑
	err := c.UpdateExpired(params.Targets, params.Expired)
	if err != nil {
		log.Printf("[ChangeExpiredController] UpdateExpired error:%+v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.WriteString("更新过期时间失败")
		return
	}

	// 返回成功响应
	ctx.ContentType("application/json")
	_, err = ctx.WriteString("{}")
	if err != nil {
		return
	}

}
