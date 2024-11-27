package impl

import (
	"github.com/kataras/iris/v12/x/errors"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
	"short-url-rw-github/src/utils"
	"strings"
	"time"
)

type GenerateService struct {
	interfaces.IValidURL
	interfaces.IDataAccessLayer
	interfaces.ILinkService
	EnvVariables config.Config
}

func (g *GenerateService) Generate(urls []string, expiredTs int64) (map[string]string, error) {

	// 设置默认过期时间
	if expiredTs == 0 {
		expiredTs = time.Now().AddDate(0, 0, 7).Unix()
	}

	results := make(map[string]string)
	for _, url := range urls {
		url = strings.TrimSpace(url)

		// 验证URL合法性
		if !g.IsValidURL(url) {
			return nil, errors.New("请提供正确的链接")
		}

		// 检查数据库是否已有记录
		existingLink, err := g.FindByOriginalUrl(url)
		if err == nil && existingLink != nil {
			results[utils.MD5Hex(url)] = g.EnvVariables.Origin + "/" + existingLink.ShortID
			continue
		}

		// 生成短链接
		shortID, err := g.generateUniqueShortID()
		if err != nil {
			return nil, err
		}

		// 保存到数据库
		link := &models.Link{
			ID:          0,
			ShortID:     shortID,
			OriginalURL: url,
			ExpiredTs:   expiredTs,
			Status:      0,
			Remark:      nil,
			CreateTime:  time.Now(),
		}
		if err := g.Create(link); err != nil {
			return nil, err
		}
		results[utils.MD5Hex(url)] = g.EnvVariables.Origin + "/" + link.ShortID
	}
	return results, nil
}

// 生成短链接并存入数据库
func (g *GenerateService) generateToDB(url string, expiredTs int64) (string, error) {
	// 检查数据库中是否已有对应的原始链接
	existingLink, err := g.FindByOriginalUrl(url)
	if err == nil && existingLink != nil {
		return existingLink.ShortID, nil
	}
	// 生成短链接ID
	shortID := utils.GenerateShortID()
	for i := 0; i < 3; i++ {
		isUsed, _ := g.CheckShortIDUsed(shortID)
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
	if err := g.Create(link); err != nil {
		return "", err
	}
	return shortID, nil
}

// 生成唯一短链接
func (g *GenerateService) generateUniqueShortID() (string, error) {
	for i := 0; i < 3; i++ {
		shortID := utils.GenerateShortID()
		isUsed, err := g.CheckShortIDUsed(shortID)
		if err != nil {
			return "", err
		}
		if !isUsed {
			return shortID, nil
		}
	}
	return "", errors.New("短链接生成冲突")
}
