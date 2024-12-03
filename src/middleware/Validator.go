package middleware

import (
	"github.com/kataras/iris/v12"
	"short-url-4go/src/config"
)

// CheckApiSecret 验证 api-secret 请求头
func CheckApiSecret() iris.Handler {
	return func(ctx iris.Context) {
		// 获取 api-secret 请求头
		secret := ctx.GetHeader("Api-Secret")

		// 验证 api-secret
		if secret == config.EnvVariables.ApiSecret {
			ctx.Next() // 验证成功，继续执行下一个处理逻辑
		} else {
			ctx.StopWithStatus(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{
				"error": "Unauthorized",
			})
		}
	}
}

// CheckToken 验证 token请求头
func CheckToken() iris.Handler {
	return func(ctx iris.Context) {
		// 获取 Token 请求头
		secret := ctx.GetHeader("Token")

		// 验证 Token
		if secret == config.EnvVariables.Token {
			ctx.Next() // 验证成功，继续执行下一个处理逻辑
		} else {
			ctx.StopWithStatus(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{
				"error": "Unauthorized",
			})
		}
	}
}
