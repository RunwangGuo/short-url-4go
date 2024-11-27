package middleware

import (
	"net/http"
	"strings"
)

const HEADER_KEY = "Api-Secret"

type ApiValidateMiddleware struct {
	secret string
}

func NewApiValidateMiddleware(secret string) *ApiValidateMiddleware {
	return &ApiValidateMiddleware{secret: secret}
}

// ServeHTTP 实现 http.Handler 接口
func (m *ApiValidateMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// 检查请求路径是否以 "/api/" 开头
	if !strings.HasPrefix(r.URL.Path, "/api/") {
		next(w, r)
		return
	}

	// 获取请求头中的 "Api-Secret" 值
	token := r.Header.Get(HEADER_KEY)

	// 验证 token 是否与预设的 secret 相同
	if token == m.secret {
		next(w, r)
	} else {
		// 如果 token 无效，则返回 401 Unauthorized 错误
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
