package controllers

import (
	"github.com/kataras/iris/v12"
	"log"
	"short-url-rw-github/src/interfaces"
	"short-url-rw-github/src/models"
)

type SearchController struct {
	interfaces.ISearchService
}

func (s *SearchController) search(ctx iris.Context) {

	// 获取查询参数
	keyword := ctx.URLParamDefault("keyword", "")
	page := ctx.URLParamIntDefault("page", 1)
	size := ctx.URLParamIntDefault("size", 30)

	// 调用服务层逻辑
	links, total, hitsMap, err := s.Search(keyword, page, size)
	if err != nil {
		log.Printf("SearchLinks error: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "failed to search links"})
		return
	}
	// 构造响应数据
	records := make([]models.SearchRecordItem, len(links))
	for i, link := range links {
		records[i] = models.SearchRecordItem{
			ID:          link.ID,
			ShortID:     link.ShortID,
			OriginalURL: link.OriginalURL,
			ExpiredTs:   link.ExpiredTs,
			Status:      link.Status,
			Remark:      link.Remark,
			CreateTime:  link.CreateTime,
			Hits:        hitsMap[link.ShortID],
		}
	}

	// 返回JSON响应
	ctx.JSON(iris.Map{
		"records": records,
		"pages":   (total + size - 1) / size, //总页数
		"size":    size,
	})
}
