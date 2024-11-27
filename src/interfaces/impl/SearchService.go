package impl

import (
	"github.com/kataras/iris/v12/x/errors"
	"short-url-rw-github/src/config"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
)

type SearchService struct {
	interfaces.IDataAccessLayer
	interfaces.ILinkService
	interfaces.IAccessLogService
	EnvVariables config.Config
}

// SearchService 查询链接及分页信息
func (s *SearchService) SearchService(keyword string, page, size int) ([]models.Link, int64, map[string]int64, error) {
	if page <= 0 || size <= 0 {
		return nil, 0, nil, errors.New("invalid pagination parameters")
	}

	// 查询链接信息
	links, total, err := s.Search(keyword, page, size)
	if err != nil {
		return nil, 0, nil, err
	}

	// 查询访问记录
	hitsMap := make(map[string]int64)
	if s.EnvVariables.AccessLog {
		shortIDs := make([]string, len(links))
		for i, link := range links {
			shortIDs[i] = link.ShortID
		}
		hitsMap, err := s.BatchQueryHits(shortIDs)
		if err != nil {
			return nil, 0, nil, err
		}
	}
	return links, total, hitsMap, nil
}
