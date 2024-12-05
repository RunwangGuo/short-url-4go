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

/*func (m *MySQLClient) Create(data interface{}) error {
	err := m.DB.Create(data).Error
	if err != nil {
		return fmt.Errorf("新增记录时发生错误: %v", err)
	}
	return nil
}*/

// CreateLink 创建Link表
func (m *MySQLClient) CreateLink(link *models.Link) error {
	err := m.DB.Create(link).Error
	if err != nil {
		return fmt.Errorf("新增记录时发生错误: %v", err)
	}
	return nil
}

// CreateAccessLog 创建AccessLog表
func (m *MySQLClient) CreateAccessLog(accessLog *models.AccessLog) error {
	err := m.DB.Create(accessLog).Error
	if err != nil {
		return fmt.Errorf("新增记录时发生错误: %v", err)
	}
	return nil
}

// UpdateLink 更新Link表
func (m *MySQLClient) UpdateLink(column string, value interface{}, query string, values ...interface{}) error {
	return m.DB.Model(&models.Link{}).
		Where(query, values...).
		Update(column, value).
		Error
}

// FindLinkByCondition 查询Link表
// TODO: 一个表一个操作模型
func (m *MySQLClient) FindLinkByCondition(condition string, value string) (*models.Link, error) {

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

/*func (m *MySQLClient) FindByCondition(condition string, value string) (*models.AccessLog, error) {

	var ac models.AccessLog
	err := m.DB.Where(condition, value).First(&ac).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果找不到记录，返回 nil 和 nil 错误
			return nil, nil
		}
		// 如果出现其他错误，返回错误
		return nil, fmt.Errorf("查找链接时发生错误: %v", err)
	}

	return &ac, nil
}*/

// CountAccessLogByCondition 统计AccessLog表
func (m *MySQLClient) CountAccessLogByCondition(condition string, value string) int64 {
	var count int64
	m.DB.Model(&models.AccessLog{}).Where(condition, value).Count(&count)
	//m.DB.Model((&User{}).Where(condition, value).Count(&count)
	return count
}

// PaginationLink 分页查询Link表
func (m *MySQLClient) PaginationLink(params *models.SearchParams) (models.PaginationResult, error) {
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
