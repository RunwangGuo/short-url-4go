package interfaces

import (
	"net/http"
	"short-url-4go/src/models"
)

type ILinkService interface {
	Redirect(shortID string, headers http.Header) (*string, error)
	Generate(url string, expiredTs int64) (string, error)
	//Generate(urls []string, expiredTs int64) (map[string]string, error)
	//FindByOriginalURL(url string) (*models.Link, error) //根据原始链接查找记录
	//CheckShortIDUsed(shortID string) (bool, error)
	//FindByShortID(shortId string) (*models.Link, error)
	Search(params *models.SearchParams) (*models.SearchResponse, error)
	UpdateStatus(targets []string, status models.LinkStatusEnum) error // 最新的
	UpdateRemark(targets []string, remark string) error                // 最新的
	UpdateExpired(targets []string, expiredTs int64) error
}
