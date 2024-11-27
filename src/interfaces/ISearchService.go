package interfaces

import "short-url-rw-github/src/models"

type ISearchService interface {
	Search(keyword string, page, size int) ([]models.Link, int, map[string]int64, error)
}
