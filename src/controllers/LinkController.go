package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"net/http"
	"short-url-4go/src/interfaces"
	"short-url-4go/src/models"
	"strings"
	"time"
)

type LinkController struct {
	interfaces.ILinkService
	Logger *zap.Logger
}

func (l *LinkController) Redirect(ctx iris.Context) {
	// 获取所有请求头
	headers := ctx.Request().Header
	headersMap := make(map[string][]string)
	for key, values := range headers {
		headersMap[key] = values
	}

	// 将请求头记录到日志
	l.Logger.Info("Request Headers", zap.Any("headers", headersMap))

	shortID := ctx.Params().Get("short_id")
	l.Logger.Info("Request ShortID", zap.String("shortID", shortID))

	// 将 headersMap 转为字符串
	var headerString strings.Builder
	for key, values := range headersMap {
		for _, value := range values {
			headerString.WriteString(key + ": " + value + "\n")
		}
	}

	// 调用服务处理重定向逻辑
	redirectURL, err := l.ILinkService.Redirect(shortID, headerString.String())
	l.Logger.Info("Redirect", zap.Any("redirectURL", redirectURL))

	switch {
	case redirectURL == "404":
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": shortID + "  源链接未找到"})
		return
	case redirectURL == "410":
		ctx.StatusCode(iris.StatusGone)
		ctx.JSON(iris.Map{"error": shortID + "  源链接被禁用"})
		return
	case redirectURL == "411":
		ctx.StatusCode(iris.StatusGone)
		ctx.JSON(iris.Map{"error": shortID + "  源链接已过期"})
		return
	}

	/*	if redirectURL == "404" {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{"error": shortID + "  源链接未找到"})
			return
		}
		if redirectURL == "410" {
			ctx.StatusCode(iris.StatusGone)
			ctx.JSON(iris.Map{"error": shortID + "  源链接被禁用"})
			return
		}

		if redirectURL == "411" {
			ctx.StatusCode(iris.StatusGone)
			ctx.JSON(iris.Map{"error": shortID + "  源链接已过期"})
			return
		}*/

	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(redirectURL, http.StatusTemporaryRedirect)
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
	// 获取所有请求头
	headers := ctx.Request().Header
	headersMap := make(map[string][]string)
	for key, values := range headers {
		headersMap[key] = values
	}

	// 将请求头记录到日志
	l.Logger.Info("Request Headers", zap.Any("headers", headersMap))

	// 定义结构体实例
	var req models.ChangeStatusReq

	// 解析请求体
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "无效的请求体" + err.Error()})
		return
	}

	// 调用service更新状态
	err := l.UpdateStatus(req.Targets, req.Status)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "状态更新失败" + err.Error()})
		return
	}

	// 返回成功响应
	ctx.JSON(iris.Map{"message": "状态更新成功"})
}

// ChangeExpired 修改过期时间的控制器
func (l *LinkController) ChangeExpired(ctx iris.Context) {
	// 获取所有请求头
	headers := ctx.Request().Header
	headersMap := make(map[string][]string)
	for key, values := range headers {
		headersMap[key] = values
	}

	// 将请求头记录到日志
	l.Logger.Info("Request Headers", zap.Any("headers", headersMap))

	// 定义结构体实例
	var req models.ChangeExpiredReq

	// 解析请求体
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 校验过期时间是否合理
	if req.Expired < time.Now().Unix() {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "请提供不小于当前日期的过期时间"})
		return
	}

	/*	// 校验时间戳是否合理
		if !utils.IsReasonableTimestamp(params.Expired) {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.WriteString("请提供不小于当前日期的过期时间")
			return
		}*/

	// 调用Service处理业务逻辑
	err := l.UpdateExpired(req.Targets, req.Expired)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "过期时间更新失败" + err.Error()})
		return
	}

	// 返回成功响应
	ctx.JSON(iris.Map{"message": "过期时间更新成功"})
}

// Remark 修改备注的控制器
func (l *LinkController) Remark(ctx iris.Context) {
	// 获取所有请求头
	headers := ctx.Request().Header
	headersMap := make(map[string][]string)
	for key, values := range headers {
		headersMap[key] = values
	}

	// 将请求头记录到日志
	l.Logger.Info("Request Headers", zap.Any("headers", headersMap))

	// 定义结构体实例
	var req models.RemarkReq

	// 解析请求体
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用Service处理逻辑
	//results, err := r.IGenerateService.Generate(params.URLs, params.ExpiredTs)
	err := l.ILinkService.UpdateRemark(req.Targets, req.Remark)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "更新备注失败: " + err.Error()})
		return
	}

	// 返回成功响应
	ctx.JSON(iris.Map{"message": "备注更新成功"})
}
