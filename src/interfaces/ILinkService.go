package interfaces

import (
	"short-url-4go/src/models"
)

type ILinkService interface {
	Generate(urls []string, expiredTs int64) (map[string]string, error)
	FindByOriginalURL(url string) (*models.Link, error) //根据原始链接查找记录
	CheckShortIDUsed(shortID string) (bool, error)
	FindByShortID(shortId string) (*models.Link, error)
	Create(data *models.Link) error
	// Search(keyword string, page, size int) ([]models.Link, int, error)
	Search(params *models.SearchParams) (*models.SearchResponse, error)
	// UpdateStatus Search(keyword string, page, size int) ([]models.Link, int, map[string]int64, error)
	// UpdateStatus(targets []string, status string) error
	UpdateStatus(targets []string, status models.LinkStatusEnum) error // 最新的
	// UpdateRemark UpdateRemark(targets []string, remark string) error
	UpdateRemark(targets []string, remark string) error // 最新的
	UpdateExpired(targets []string, expiredTs int64) error
	GetRedirectURL(shortID string) (string, string, error) // 返回URL、模板、错误
}
