package models

import "time"

type AccessLog struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ShortID    string    `json:"short_id" gorm:"size:50;not null;comment:'短链ID'"`
	ReqHeaders string    `json:"req_headers" gorm:"type:longtext;comment:'请求头'"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;not null;comment:'创建时间'"`
}
