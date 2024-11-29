package interfaces

import (
	"short-url-rw-github/src/models"
)

type IDataAccessLayer interface {
	Create(data models.Link) (*models.Link, error)
	Update(column string, value interface{}, query string, values ...interface{}) error
	FindByCondition(condition string, value string) (*models.Link, error)
}
