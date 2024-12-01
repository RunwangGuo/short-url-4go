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

// SearchParams 分页请求参数
type SearchParams struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"` // 当前页码
	Size    int    `json:"size"` // 每页大小
}

// PaginationResult 分页结果结构体
type PaginationResult struct {
	Records []Link `json:"records"` // 当前页的记录
	Pages   int    `json:"pages"`   // 总页数
}

// SearchResponse 搜索响应
type SearchResponse struct {
	Records []SearchRecordItem `json:"records"`
	Pages   int                `json:"pages"`
	Size    int                `json:"size"`
}

// SearchRecordItem 响应记录项
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

type ChangeStatusReq struct {
	Targets []string `json:"s"`
	//Status  LinkStatusEnum `json:"status"`
	Status LinkStatusEnum `json:"status"`
}

type ChangeExpiredReq struct {
	Targets []string `json:"targets"` // 短链接的ID列表
	Expired int64    `json:"expired"` // 新的过期时间戳
}

type RemarkReq struct {
	Targets []string `json:"targets"` //短链接ID列表
	Remark  string   `json:"remark"`  // 新的备注内容
}
