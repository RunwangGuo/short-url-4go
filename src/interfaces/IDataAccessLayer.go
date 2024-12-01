package interfaces

import (
	"short-url-4go/src/models"
)

type IDataAccessLayer interface {
	Create(data models.Link) (*models.Link, error)
	Update(column string, value interface{}, query string, values ...interface{}) error
	FindByCondition(condition string, value string) (*models.Link, error)
	Pagination(params *models.SearchParams) (models.PaginationResult, error)
	CountByCondition(table interface{}, condition string, value string) int64
}
