package infrastrctures

import (
	"fmt"
	"github.com/kataras/iris/v12/x/errors"
	"gorm.io/gorm"
	"short-url-rw-github/src/models"
)

type MySQLClient struct {
	DB *gorm.DB
}

func (m *MySQLClient) Create(data models.Link) (*models.Link, error) {
	link := &models.Link{
		ShortID:     data.ShortID,
		OriginalURL: data.OriginalURL,
		ExpiredTs:   data.ExpiredTs,
		Status:      data.Status,
	}

	err := m.DB.Create(link).Error
	if err != nil {
		return nil, fmt.Errorf("新增记录时发生错误: %v", err)
	}

	return link, nil
}

// Update 通用更新方法
func (m *MySQLClient) Update(column string, value interface{}, query string, values ...interface{}) error {
	return m.DB.Model(&models.Link{}).
		Where(query, values...).
		Update(column, value).
		Error
}

// FindByCondition 通用条件查询
func (m *MySQLClient) FindByCondition(condition string, value string) (*models.Link, error) {

	err := m.DB.Where(condition, value).First(&models.Link{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果找不到记录，返回 nil 和 nil 错误
			return nil, nil
		}
		// 如果出現其他錯誤，返回錯誤
		return nil, fmt.Errorf("查找链接时发生错误: %v", err)
	}

	return &models.Link{}, nil
}
