package interfaces

import "short-url-rw-github/src/models"

type IDataAccessLayer interface {
	//Create(data *models.Link) error
	FindByOriginalUrl(originalUrl string) (*models.Link, error) //根据原始链接查找记录
}
