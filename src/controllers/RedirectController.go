package controllers

import (
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"net/http"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/utils"
)

type RedirectController struct {
	interfaces.IAccessLogService
	interfaces.ILinkService
	zap *zap.Logger
}

func (r *RedirectController) Redirect(ctx iris.Context) {
	shortID := ctx.Params().Get("short_id")

	// 如果ID中存在其他字符，剔除
	if len(shortID) > utils.ShortIDLen {
		shortID = shortID[:utils.ShortIDLen]
	}

	// 异步记录访问日志
	go func() {
		if err := r.RecordAccessLog(shortID, ctx.Request().Header); err != nil {
			r.zap.Error("Failed to record access log", zap.Error(err))
		}
	}()

	// 获取重定向URL
	redirectURL, template, err := r.GetRedirectURL(shortID)
	if err != nil {
		r.zap.Error("Failed to get redirect URL", zap.String("short_id", shortID), zap.Error(err))
	}

	// 根据结果返回响应
	if redirectURL == "" {
		switch template {
		case "error/404.html":
			ctx.NotFound()
		case "disabled.html":
			ctx.StatusCode(http.StatusForbidden)
			ctx.View("disabled.html")
		case "expired.html":
			ctx.StatusCode(http.StatusGone)
			ctx.View("expired.html")
		}
		return
	}
	ctx.Redirect(redirectURL, http.StatusTemporaryRedirect)
}
