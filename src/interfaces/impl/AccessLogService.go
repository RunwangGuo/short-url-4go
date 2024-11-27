package impl

import (
	"gorm.io/gorm"
	"net/http"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
	"time"
)

// AccessLogService 提供对 AccessLog 的操作
type AccessLogService struct {
	DB *gorm.DB
	interfaces.IDataAccessLayer
	EnvVariables *config.Config
}

// Add 方法：将数据插入到数据库中
func (a *AccessLogService) Add(db *gorm.DB, data models.AccessLog) (*models.AccessLog, error) {
	// 使用Create方法插入数据
	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// RecordAccessLog 日志写入逻辑
func (a *AccessLogService) RecordAccessLog(shortID string, headers http.Header) error {
	// 将 HTTP Header 格式化为字符串
	reqHeaders := ""
	for key, values := range headers {
		for _, value := range values {
			reqHeaders += key + ": " + value + "\n"
		}
	}

	// 创建日志模型
	accessLog := models.AccessLog{
		ShortID:    shortID,
		ReqHeaders: reqHeaders,
		CreateTime: time.Now(),
	}

	// 写入数据库
	a.DB.Exec(`
		INSERT INTO access_log (short_id, req_headers, create_time)
		VALUES (?,?,?)`, shortID, reqHeaders, accessLog.CreateTime)
	return nil
}

// BatchQueryHits 批量查询点击次数
func (a *AccessLogService) BatchQueryHits(shortIDs []string) (map[string]int64, error) {
	hitsMap := make(map[string]int64)
	if len(shortIDs) == 0 {
		return hitsMap, nil
	}

	rows, err := a.DB.Table("access_logs").
		Select("short_id, COUNT(*) AS hits").
		Where("short_id IN ?", shortIDs).
		Group("short_id").
		Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var shortID string
		var hits int64
		if err := rows.Scan(&shortID, &hits); err != nil {
			return nil, err
		}
		hitsMap[shortID] = hits
	}
	return hitsMap, nil
}
