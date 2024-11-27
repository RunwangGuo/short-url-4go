package models

import "time"

const (
	LinkStatusEnabled  = 1
	LinkStatusDisabled = 2
)

type AccessLog struct {
	ID         uint64    `json:"id" gorm:"primaryKey;column:id"`
	ShortID    string    `json:"short_id" gorm:"column:short_id"`
	ReqHeaders string    `json:"req_headers" gorm:"column:req_headers"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (AccessLog) TableName() string {
	return "access_log"
}
