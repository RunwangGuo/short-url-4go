package infrastrctures

import (
	"gorm.io/gorm"
	"log"
	"short-url-rw-github/src/models"
)

type MySQLClient struct {
	DB *gorm.DB
}

func (m *MySQLClient) Create(data *models.Link) error {
	if err := m.DB.Create(&data).Error; err != nil {
		log.Printf("[link service] create error: %v", err)
		return err
	}
	return nil
}

func (m *MySQLClient) FindByOriginalURL(url string) (*models.Link, error) {
	var link models.Link
	err := m.DB.Where("original_url = ?", url).First(&link).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("[data access layer] no record found for original_url: %s", url)
			return nil, nil // 没有找到记录，返回 nil
		}
		log.Printf("[data access layer] find_by_original_url error for url %s: %v", url, err)
		return nil, err // 数据库查询出错
	}
	return &link, nil
}
