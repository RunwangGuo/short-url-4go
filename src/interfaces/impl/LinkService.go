package impl

import (
	"github.com/kataras/iris/v12/x/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
	"time"
)

type LinkService struct {
	DB *gorm.DB
	interfaces.IDataAccessLayer
	interfaces.ICacheLayer
	EnvVariables *config.Config
	zap          *zap.Logger
}

// FindByOriginalURL 根据原始链接查找记录
//func (l *LinkService) FindByOriginalURL(url string) (*models.Link, error)

func (l *LinkService) GetRedirectURL(shortID string) (string, string, error) {
	/*	// 1 查询缓存
		if url, err := l.Get(shortID);{
			if url != "" {
				l.zap.Info("Cache hit", zap.String("short_id", shortID), zap.String("url", url))
				return url, "", nil
			}
			l.zap.Warn("Cache hit but URL is invalid", zap.String("short_id", shortID))
			return "", "error/404.html", nil
		}*/

	// 1. 查询缓存
	value, err := l.Get(shortID)
	if err != nil {
		l.zap.Error("Cache error", zap.String("short_id", shortID), zap.Error(err))
		return value, "", nil
	}

	// 2 查询数据库
	link, err := l.FindByShortID(shortID)
	if err != nil {
		l.zap.Error("Database error", zap.String("short_id", shortID), zap.Error(err))
		return "", "", err
	}

	if link == nil {
		l.zap.Warn("short ID not found", zap.String("short_id", shortID))
		l.Set(shortID, "")
		return "", "error/404.html", nil
	}

	// 3 检查链接状态
	if link.Status == models.LinkStatusDisabled {
		l.zap.Info("Link disabled", zap.String("short_id", shortID))
		l.Set(shortID, "")
		return "", "disabled.html", nil
	}

	if link.ExpiredTs > 0 && link.ExpiredTs < time.Now().UnixMilli() {
		l.zap.Info("Link expired", zap.String("short_id", shortID))
		l.Set(shortID, "")
		return "", "expired.html", nil
	}

	// 4 缓存结果并返回
	l.Set(shortID, link.OriginalURL)
	l.zap.Info("Redirect URL found", zap.String("short_id", shortID), zap.String("url", link.OriginalURL))
	return link.OriginalURL, "", nil
}

// CheckShortIDUsed 检查 ShortID 是否已被使用
func (l *LinkService) CheckShortIDUsed(shortID string) (bool, error) {
	var count int64
	if err := l.DB.Model(&models.Link{}).Where("short_id = ?", shortID).Count(&count).Error; err != nil {
		log.Printf("[link service] check_short_id_used: %s error: %v", shortID, err)
		return false, err
	}
	return count > 0, nil
}

// FindByShortID 根据ShortID查找记录
func (l *LinkService) FindByShortID(shortId string) (*models.Link, error) {
	var link models.Link
	if err := l.DB.Where("short_id = ?", shortId).First(&link).Error; err != nil {
		log.Printf("[link service] find_by_short_id: %s error: %v", shortId, err)
		return nil, err
	}
	return &link, nil
}

// Create 创建记录
func (l *LinkService) Create(data *models.Link) error {
	if err := l.DB.Create(&data).Error; err != nil {
		log.Printf("[link service] create error: %v", err)
		return err
	}
	return nil
}

// Search 根据关键字和分页条件查询链接
func (l *LinkService) Search(keyword string, page, size int) ([]models.Link, int, error) {
	var links []models.Link
	var total int64

	query := l.DB.Model(&models.Link{})
	if keyword != "" {
		query = query.Where("original_url LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err = query.Order(offset).Limit(size).Find(&links).Error

	if err != nil {
		return nil, 0, err
	}

	return links, int(total), nil
}

//// UpdateStatus 更新状态
//func (l *LinkService) UpdateStatus(targets []string, status string) error {
//	if len(targets) == 0 {
//		return nil
//	}
//	return l.DB.Model(&models.Link{}).
//		Where("short_id = ?", targets).
//		Update("status", status).Error
//}

// UpdateStatus 批量更新状态
func (l *LinkService) UpdateStatus(targets []string, status string) error {
	if len(targets) == 0 {
		return errors.New("targets cannot be empty")
	}

	// 更新数据库状态
	err := l.UpdateStatus(targets, status)
	if err != nil {
		log.Printf("[LinkService] UpdateStatus error: %v", err)
		return err
	}

	// 清除缓存中的相关条目
	err = l.Remove(targets)
	if err != nil {
		log.Printf("[LinkService] RemoveLink error: %v", err)
		return err
	}
	return nil
}

// UpdateExpired 批量更新过期时间
func (l *LinkService) UpdateExpired(targets []string, expiredTs int64) error {

	// 更新数据库中的过期时间
	if err := l.DB.Model(&models.Link{}).
		Where("short_id IN ?", targets).
		Update("expired_ts", expiredTs).Error; err != nil {
		log.Printf("[link service] update_expired error: %v", err)
		return err
	}

	// 清除缓存中的相关条目
	err := l.Remove(targets)
	if err != nil {
		log.Printf("[LinkService] RemoveLink error: %v", err)
		return err
	}

	return nil
}

// UpdateRemark 更新数据库中的备注
func (l *LinkService) UpdateRemark(targets []string, remark string) error {
	if err := l.DB.Model(&models.Link{}).
		Where("short_id IN ?", targets).
		Update("remark", remark).Error; err != nil {
		log.Printf("[link service] update_remark error: %v", err)
		return err
	}
	return nil
}
