package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const ShortIDLen = 5

// GenerateShortID 生成短链接ID
func GenerateShortID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, ShortIDLen)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// IsValidURL 检查是否是有效的 URL
func IsValidURL(str string) bool {
	_, err := url.ParseRequestURI(str)
	return err == nil
}

// Hash256Hex 计算 SHA-256 哈希值（十六进制）
func Hash256Hex(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Hash1Hex 计算 SHA-1 哈希值（十六进制）
func Hash1Hex(data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// MD5Hex 计算 MD5 哈希值（十六进制）
func MD5Hex(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// ValidateHeaderToken 验证 Header 中的 Token
func ValidateHeaderToken(headerValue string, token string) bool {
	return strings.TrimSpace(headerValue) == token
}

// IsReasonableTimestamp 检查时间戳是否合理
func IsReasonableTimestamp(value int64) bool {
	if value == 0 {
		return true
	}
	timestamp := time.Unix(value, 0)
	now := time.Now().UTC()
	return !timestamp.IsZero() && timestamp.After(now)
}
