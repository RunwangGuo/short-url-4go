package services

import (
	"fmt"
	"github.com/kataras/iris/v12/x/errors"
	"go.uber.org/zap"
	"short-url-4go/src/config"
	"short-url-4go/src/interfaces"
	"short-url-4go/src/models"
	"short-url-4go/src/utils"
	"time"
)

type LinkService struct {
	interfaces.IDataAccessLayer
	interfaces.ICacheLayer
	Logger *zap.Logger
}

/*// FindByOriginalURL 根据原始链接查找记录
func (l *LinkService) FindByOriginalURL(url string) (*models.Link, error) {
	var link models.Link
	err := l.DB.Where("original_url = ?", url).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("[data access layer] no record found for original_url: %s", url)
			return nil, nil // 没有找到记录，返回 nil
		}
		log.Printf("[data access layer] find_by_original_url error for url %s: %v", url, err)
		return nil, err // 数据库查询出错
	}
	return &link, nil
}*/

/*// FindByOriginalURL 根据原始链接查找记录
func (l *LinkService) FindByOriginalURL(url string) (*models.Link, error) {
	record, err := l.FindLinkByCondition("original_url = ?", url)
	if err != nil {
		return nil, err
	}
	return record, nil
}*/

/*// FindByShortID 根据ShortID查找记录
func (l *LinkService) FindByShortID(shortId string) (*models.Link, error) {
	var link models.Link
	if err := l.DB.Where("short_id = ?", shortId).First(&link).Error; err != nil {
		log.Printf("[link service] find_by_short_id: %s error: %v", shortId, err)
		return nil, err
	}
	return &link, nil
}*/

/*// FindByShortID 根据ShortID查找记录
func (l *LinkService) FindByShortID(shortId string) (*models.Link, error) {
	record, err := l.FindLinkByCondition("short_id = ?", shortId)
	if err != nil {
		return nil, err
	}
	return record, nil
}*/

// CheckShortIDUsed 检查 ShortID 是否已被使用
/*func (l *LinkService) CheckShortIDUsed(shortID string) (bool, error) {
	var count int64
	if err := l.DB.Model(&models.Link{}).Where("short_id = ?", shortID).Count(&count).Error; err != nil {
		log.Printf("[link service] check_short_id_used: %s error: %v", shortID, err)
		return false, err
	}
	return count > 0, nil
}*/

/*// CheckShortIDUsed 检查 ShortID 是否已被使用
func (l *LinkService) CheckShortIDUsed(shortId string) (bool, error) {
	record, err := l.FindLinkByCondition("short_id = ?", shortId)
	if err != nil {
		return false, err
	}
	if record == nil {
		return false, nil
	}
	return true, nil
}*/

/*// Create 创建记录
func (l *LinkService) Create(data *models.Link) error {
	if err := l.DB.Create(&data).Error; err != nil {
		log.Printf("[link service] create error: %v", err)
		return err
	}
	return nil
}*/

/*// Create 创建记录
func (l *LinkService) Create(data *models.Link) error {
	err := l.Create(data)
	if err != nil {
		return err
	}
	return nil
}*/

func (l *LinkService) Generate(url string, expiredTs int64) (string, error) {

	// 检查是否存在相同的原始链接
	record, err := l.FindLinkByCondition("original_url = ?", url)
	if err != nil && record != nil && record.ShortID != "" {
		return record.ShortID, nil
	}

	// 生成短链接ID, 最多尝试3次, 防碰撞
	var shortID string
	for i := 0; i < 3; i++ {
		shortID = utils.GenerateShortID()
		used, _ := l.FindLinkByCondition("short_id = ?", shortID)
		if used == nil {
			break
		}
	}

	// 保存短链接到数据库
	link := &models.Link{
		//ID:          0,
		ShortID:     shortID,
		OriginalURL: url,
		ExpiredTs:   expiredTs,
		Status:      0,
		Remark:      nil,
		CreateTime:  time.Now(),
	}
	err = l.CreateLink(link)
	if err != nil {
		return "", err
	}

	return shortID, nil
}

/*// 生成短链接并存入数据库
func (l *LinkService) generateToDB(url string, expiredTs int64) (string, error) {
	// 检查数据库中是否已有对应的原始链接
	existingLink, err := l.FindByOriginalURL(url)
	if err == nil && existingLink != nil {
		return existingLink.ShortID, nil
	}
	// 生成短链接ID
	shortID := utils.GenerateShortID()
	for i := 0; i < 3; i++ {
		isUsed, _ := l.CheckShortIDUsed(shortID)
		if isUsed {
			shortID = utils.GenerateShortID()
		} else {
			break
		}
		if i == 2 {
			return "", errors.New("短链接生成冲突")
		}
	}

	var link *models.Link
	// 保存到数据库
	if err := l.Create(link); err != nil {
		return "", err
	}
	return shortID, nil
}*/

/*// 生成唯一短链接
func (l *LinkService) generateUniqueShortID() (string, error) {
	for i := 0; i < 3; i++ {
		shortID := utils.GenerateShortID()
		isUsed, err := l.CheckShortIDUsed(shortID)
		if err != nil {
			return "", err
		}
		if !isUsed {
			return shortID, nil
		}
	}
	return "", errors.New("短链接生成冲突")
}*/

/*// SearchService 查询链接及分页信息
func (l *LinkService) SearchService(keyword string, page, size int) ([]models.Link, int64, map[string]int64, error) {
	if page <= 0 || size <= 0 {
		return nil, 0, nil, errors.New("invalid pagination parameters")
	}

	// 查询链接信息
	links, total, err := l.Search(keyword, page, size)
	if err != nil {
		return nil, 0, nil, err
	}

	// 查询访问记录
	hitsMap := make(map[string]int64)
	if config.EnvVariables.AccessLog {
		shortIDs := make([]string, len(links))
		for i, link := range links {
			shortIDs[i] = link.ShortID
		}
		hitsMap, err := l.BatchQueryHits(shortIDs)
		if err != nil {
			return nil, 0, nil, err
		}
	}
	return links, total, hitsMap, nil
}*/

/*func (l *LinkService) GetRedirectURL(shortID string) (*string, string, error) {
		// 1 查询缓存
		//if url, err := l.Get(shortID);{
		//	if url != "" {
		//		l.zap.Info("Cache hit", zap.String("short_id", shortID), zap.String("url", url))
		//		return url, "", nil
		//	}
		//	l.zap.Warn("Cache hit but URL is invalid", zap.String("short_id", shortID))
		//	return "", "error/404.html", nil
		//}

	// 1. 查询缓存
	value, err := l.Get(shortID)
	if err != nil {
		l.Logger.Error("Cache error", zap.String("short_id", shortID), zap.Error(err))
		return value, "", nil
	}

	// 2 查询数据库
	link, err := l.FindByShortID(shortID)
	if err != nil {
		l.Logger.Error("Database error", zap.String("short_id", shortID), zap.Error(err))
		return "", "", err
	}

	if link == nil {
		l.Logger.Warn("short ID not found", zap.String("short_id", shortID))
		l.Set(shortID, "")
		return "", "error/404.html", nil
	}

	// 3 检查链接状态
	if link.Status == models.LinkStatusDisabled {
		l.Logger.Info("Link disabled", zap.String("short_id", shortID))
		l.Set(shortID, "")
		return "", "disabled.html", nil
	}

	if link.ExpiredTs > 0 && link.ExpiredTs < time.Now().UnixMilli() {
		l.Logger.Info("Link expired", zap.String("short_id", shortID))
		l.Set(shortID, "")
		return "", "expired.html", nil
	}

	// 4 缓存结果并返回
	l.Set(shortID, link.OriginalURL)
	l.Logger.Info("Redirect URL found", zap.String("short_id", shortID), zap.String("url", link.OriginalURL))
	return link.OriginalURL, "", nil
}*/

/*// Search 根据关键字和分页条件查询链接
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
}*/

func (l *LinkService) Redirect(shortID string, headers string) (string, error) {
	// 如果启用了访问日志功能，记录请求头信息
	if config.EnvVariables.AccessLog {
		go func(shortID string, headers string) {
			/*			// 将请求头转为字符串
						var headerString strings.Builder
						for key, values := range headers {
							for _, value := range values {
								headerString.WriteString(key + ": " + value + "\n")
							}
						}*/
			// 记录访问日志
			accessLog := &models.AccessLog{
				ShortID:    shortID,
				ReqHeaders: headers,
				CreateTime: time.Now(),
			}
			l.Logger.Info("Creating access log", zap.Any("accessLog", accessLog))
			if err := l.CreateAccessLog(accessLog); err != nil {
				l.Logger.Error("Failed to add access log")
			}
		}(shortID, headers)
	}

	// 检查缓存是否存在
	cache, err := l.Get(shortID)
	if cache == "" || err != nil {
		l.Logger.Info("获得缓存失败", zap.String("cache是", cache), zap.Error(err))

	}
	if err != nil && cache != "" {
		l.Logger.Info("缓存中存在" + shortID + "对应的OriginalURL，值是" + cache)
		return cache, nil
	}

	// 查询数据库，获取短链接对应的原始链接
	record, err := l.FindLinkByCondition("short_id = ?", shortID)
	if err != nil {
		return "", err
	}

	if record == nil {
		return "404", err
	}

	// 检查链接是否已经被禁用
	if record.Status == models.Disabled {
		return "410", errors.New("link is disabled")
	}

	// 检查链接是否已经过期
	if record.ExpiredTs > 0 && record.ExpiredTs < time.Now().Unix() {
		return "411", errors.New("link is expired")
	}

	l.Logger.Info("", zap.Int64("记录的时间是", record.ExpiredTs), zap.Int64("", time.Now().Unix()))

	// 缓存并返回原始链接
	l.Logger.Info("shortID是" + shortID)
	l.Logger.Info("OriginalURL是" + record.OriginalURL)
	_ = l.Set(shortID, record.OriginalURL)
	cache, _ = l.Get(shortID)
	l.Logger.Info("缓存中" + shortID + "对应的OriginalURL是" + cache)
	return record.OriginalURL, nil
}

// Search Search处理逻辑
func (l *LinkService) Search(params *models.SearchParams) (*models.SearchResponse, error) {
	// 获取分页数据
	paginationResult, err := l.PaginationLink(params)
	fmt.Print(paginationResult)
	if err != nil {
		return nil, err
	}

	// 初始化访问次数
	hitsMap := make(map[string]int64)
	if config.EnvVariables.AccessLog {
		// 如果启用了访问日志功能，获取访问次数
		for _, link := range paginationResult.Records {
			// 获取访问次数
			hits := l.CountAccessLogByCondition("short_id = ?", link.ShortID)
			hitsMap[link.ShortID] = hits
		}
	}

	// 构造响应数据
	records := make([]models.SearchRecordItem, len(paginationResult.Records))
	for i, link := range paginationResult.Records {
		hits := hitsMap[link.ShortID] // 从 hitsMap 获取访问次数
		records[i] = models.SearchRecordItem{
			ID:          link.ID,
			ShortID:     link.ShortID,
			OriginalURL: link.OriginalURL,
			ExpiredTs:   link.ExpiredTs,
			Status:      link.Status,
			Remark:      link.Remark,
			CreateTime:  link.CreateTime,
			Hits:        hits,
		}
	}
	// 返回最终响应结果
	return &models.SearchResponse{
		Records: records,
		Pages:   paginationResult.Pages,
		Size:    params.Size,
	}, nil

}

/*// UpdateStatus 更新状态
func (l *LinkService) UpdateStatus(targets []string, status string) error {
	if len(targets) == 0 {
		return nil
	}
	return l.DB.Model(&models.Link{}).
		Where("short_id = ?", targets).
		Update("status", status).Error
}*/

/*// UpdateStatus 批量更新状态
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
}*/

func (l *LinkService) UpdateStatus(targets []string, status models.LinkStatusEnum) error {
	return l.UpdateLink("status", status, "short_id IN ?", targets)
}

/*func (l *LinkService) UpdateRemark (targets []string, remark string) error {
	if err := l.DB.Model(&models.Link{}).
		Where("short_id IN ?", targets).
		Update("remark", remark).Error; err != nil {
		log.Printf("[link service] update_remark error: %v", err)
		return err
	}
	return nil
}
*/

func (l *LinkService) UpdateRemark(targets []string, remark string) error {
	return l.UpdateLink("remark", remark, "short_id IN ?", targets)
}

/*// UpdateExpired 批量更新过期时间
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
}*/

// UpdateExpired 批量更新过期时间
func (l *LinkService) UpdateExpired(targets []string, expiredTs int64) error {
	return l.UpdateLink("expired_ts", expiredTs, "short_id IN ?", targets)
}
