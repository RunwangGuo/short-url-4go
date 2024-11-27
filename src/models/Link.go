package models

import "time"

type LinkStatusEnum uint16

const (
	Normal   LinkStatusEnum = 0
	Disabled LinkStatusEnum = 1
)

type Link struct {
	ID          uint64    `json:"id" gorm:"primaryKey;column:id"`
	ShortID     string    `json:"short_id" gorm:"column:short_id"`
	OriginalURL string    `json:"original_url" gorm:"column:original_url"`
	ExpiredTs   int64     `json:"expired_ts" gorm:"column:expired_ts"`
	Status      int16     `json:"status" gorm:"column:status"`
	Remark      *string   `json:"remark" gorm:"column:remark"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}

type GenerateReq struct {
	URLs      []string `json:"urls"`
	ExpiredTs int64    `json:"expiredTs"`
}

type SearchParams struct {
	Keyword *string `json:"keyword"`
	Page    *uint64 `json:"page"`
	Size    *uint64 `json:"size"`
}

type ChangeStatusReq struct {
	Targets []string `json:"s"`
	//Status  LinkStatusEnum `json:"status"`
	Status string `json:"status"`
}

type ChangeExpiredReq struct {
	Targets []string `json:"targets"` // 短链接的ID列表
	Expired int64    `json:"expired"` // 新的过期时间戳
}

type RemarkReq struct {
	Targets []string `json:"targets"` //短链接ID列表
	Remark  string   `json:"remark"`  // 新的备注内容
}

type SearchRecordItem struct {
	ID          uint64    `json:"id"`
	ShortID     string    `json:"shortId"`
	OriginalURL string    `json:"originalUrl"`
	ExpiredTs   int64     `json:"expiredTs"`
	Status      int16     `json:"status"`
	Remark      *string   `json:"remark"`
	CreateTime  time.Time `json:"createTime"`
	Hits        int64     `json:"hits"`
}
