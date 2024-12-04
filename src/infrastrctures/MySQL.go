package infrastrctures

import (
	"fmt"
	"github.com/kataras/iris/v12/x/errors"
	"gorm.io/gorm"
	"math"
	"short-url-4go/src/models"
)

type MySQLClient struct {
	DB *gorm.DB
}

/*func (m *MySQLClient) Create(data models.Link) error {
	link := &models.Link{
		ShortID:     data.ShortID,
		OriginalURL: data.OriginalURL,
		ExpiredTs:   data.ExpiredTs,
		Status:      data.Status,
	}

	err := m.DB.Create(link).Error
	if err != nil {
		return fmt.Errorf("新增记录时发生错误: %v", err)
	}

	return nil
}

func (m *MySQLClient) Create(data models.AccessLog) error {
	accessLog := &models.AccessLog{
		ShortID:    data.ShortID,
		ReqHeaders: data.ReqHeaders,
		CreateTime: data.CreateTime,
	}

	err := m.DB.Create(accessLog).Error
	if err != nil {
		return fmt.Errorf("新增记录时发生错误: %v", err)
	}

	return nil
}*/

func (m *MySQLClient) Create(data interface{}) error {
	err := m.DB.Create(data).Error
	if err != nil {
		return fmt.Errorf("新增记录时发生错误: %v", err)
	}
	return nil
}

// Update 通用更新方法
func (m *MySQLClient) Update(model interface{}, column string, value interface{}, query string, values ...interface{}) error {
	return m.DB.Model(model).
		Where(query, values...).
		Update(column, value).
		Error
}

// FindByCondition 通用条件查询
func (m *MySQLClient) FindByCondition(condition string, value string) (*models.Link, error) {

	var link models.Link
	err := m.DB.Where(condition, value).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果找不到记录，返回 nil 和 nil 错误
			return nil, nil
		}
		// 如果出现其他错误，返回错误
		return nil, fmt.Errorf("查找链接时发生错误: %v", err)
	}

	return &link, nil
}

// CountByCondition 通用条件计数
func (m *MySQLClient) CountByCondition(table interface{}, condition string, value string) int64 {
	var count int64
	m.DB.Model(&table).Where(condition, value).Count(&count)
	//m.DB.Model((&User{}).Where(condition, value).Count(&count)
	return count
}

// Pagination 获取分页数据
func (m *MySQLClient) Pagination(params *models.SearchParams) (models.PaginationResult, error) {
	// 初始化
	query := m.DB.Model(&models.Link{})
	if params.Keyword != "" {
		keyword := "%" + params.Keyword + "%"
		print("查询条件是" + keyword)
		query = query.Where("short_id LIKE ? OR original_url LIKE ?", keyword, keyword)
	}

	// 获取总记录数
	var totalRows int64
	query.Count(&totalRows)

	// 计算分页参数
	page := params.Page
	size := params.Size
	if page < 1 {
		page = 1
	}
	// 默认每页30条
	if size < 1 {
		size = 30
	}

	offset := (page - 1) * size
	totalPages := int(math.Ceil(float64(totalRows) / float64(size)))

	// 获取当前页数据
	var links []models.Link
	query.Order("id desc").Limit(size).Offset(offset).Find(&links)

	// 返回分页结果
	return models.PaginationResult{
		Records: links,
		Pages:   totalPages,
	}, nil
}
