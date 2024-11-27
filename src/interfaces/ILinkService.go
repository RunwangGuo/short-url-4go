package interfaces

import (
	"short-url-rw-github/src/models"
)

type ILinkService interface {
	//FindByOriginalURL(db *gorm.DB, url string) (*models.Link, error)
	CheckShortIDUsed(shortID string) (bool, error)
	FindByShortID(shortId string) (*models.Link, error)
	Create(data *models.Link) error
	Search(keyword string, page, size int) ([]models.Link, int64, error)
	//UpdateStatus(targets []string, status int16) error
	UpdateStatus(targets []string, status string) error
	UpdateExpired(targets []string, expiredTs int64) error
	UpdateRemark(targets []string, remark string) error
	GetRedirectURL(shortID string) (string, string, error) // 返回URL、模板、错误
}
