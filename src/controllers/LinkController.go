package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"log"
	"net/http"
	"short-url-4go/src/config"
	"short-url-4go/src/interfaces"
	"short-url-4go/src/models"
	"time"
)

type LinkController struct {
	interfaces.ILinkService
	Logger *zap.Logger
}

func (l *LinkController) Redirect(ctx iris.Context) {
	shortID := ctx.Params().Get("short_id")

	// 调用服务处理重定向逻辑
	redirectURL, err := l.ILinkService.Redirect(shortID, ctx.Request().Header)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, err)
	}
	ctx.Redirect(*redirectURL, http.StatusTemporaryRedirect)
}

func (l *LinkController) Generate(ctx iris.Context) {
	// 获取所有请求头
	headers := ctx.Request().Header
	headersMap := make(map[string][]string)
	for key, values := range headers {
		headersMap[key] = values
	}

	// 将请求头记录到日志
	l.Logger.Info("Request Headers", zap.Any("headers", headersMap))

	// 定义结构体实例
	var params models.GenerateReq

	// 解析请求体
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "无效的请求体" + err.Error()})
	}

	// 检查URLS是否为空
	if len(params.URLs) == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "请求体中必须包含 URLs"})
		return
	}

	// 使用解析出的数据
	// 调用Service处理逻辑
	results := make(map[string]string)
	for _, url := range params.URLs {
		l.Logger.Info("请求生成短链的长链接是" + url)
		// 调用服务层生成短链接
		shortLink, err := l.ILinkService.Generate(url, params.ExpiredTs)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		// 生成md5 hash
		hash := md5.Sum([]byte(url))
		results[hex.EncodeToString(hash[:])] = shortLink
	}

	// 序列化并返回JSON响应
	resp, err := json.Marshal(results)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("序列化响应失败")
		return
	}
	// 设置响应类型为JSON
	ctx.ContentType("application/json")
	ctx.Write(resp)
}

func (l *LinkController) Search(ctx iris.Context) {
	// 获取所有请求头
	headers := ctx.Request().Header
	headersMap := make(map[string][]string)
	for key, values := range headers {
		headersMap[key] = values
	}

	// 将请求头记录到日志
	l.Logger.Info("Request Headers", zap.Any("headers", headersMap))

	// 获取查询参数
	keyword := ctx.URLParamDefault("keyword", "")
	page := ctx.URLParamIntDefault("page", 1)
	size := ctx.URLParamIntDefault("size", 30)

	// 构造查询参数
	params := &models.SearchParams{
		Keyword: keyword,
		Page:    page,
		Size:    size,
	}

	// 调用服务层逻辑
	result, err := l.ILinkService.Search(params)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	// 返回JSON响应
	ctx.JSON(result)
}

func (l *LinkController) ChangeStatus(ctx iris.Context) {
	// 验证Token
	headerToken := ctx.GetHeader("Authorization")
	if headerToken != config.EnvVariables.Token {
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
	err := l.UpdateStatus(req.Targets, req.Status)
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

// ChangeExpired 修改过期时间的控制器
func (l *LinkController) ChangeExpired(ctx iris.Context) {

	// 校验Token
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" || headerToken != config.EnvVariables.Token {
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
	err := l.UpdateExpired(params.Targets, params.Expired)
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

// Remark 修改备注的控制器
func (l *LinkController) Remark(ctx iris.Context) {
	// 校验 Token
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" || headerToken != config.EnvVariables.Token {
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
	err := l.ILinkService.UpdateRemark(params.Targets, params.Remark)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("更新备注失败")
		return
	}

	// 返回成功响应
	ctx.ContentType("application/json")
	ctx.WriteString("{}")
}
