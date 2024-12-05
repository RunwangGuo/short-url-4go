package interfaces

import (
	"short-url-4go/src/models"
)

type IDataAccessLayer interface {
	// Create(data models.Link) (*models.Link, error)
	// Create(data interface{}) error
	CreateLink(link *models.Link) error
	CreateAccessLog(accessLog *models.AccessLog) error
	// Update(model interface{}, column string, value interface{}, query string, values ...interface{}) error
	UpdateLink(column string, value interface{}, query string, values ...interface{}) error
	// FindByCondition(condition string, value string) (*models.Link, error)
	FindLinkByCondition(condition string, value string) (*models.Link, error)
	// Pagination(params *models.SearchParams) (models.PaginationResult, error)
	PaginationLink(params *models.SearchParams) (models.PaginationResult, error)
	// CountByCondition(table interface{}, condition string, value string) int64
	CountAccessLogByCondition(condition string, value string) int64
}
