package channels

import "github.com/smartwalle/wxm"

// GetLiveListRequest 获取直播大屏直播列表请求参数
type GetLiveListRequest struct {
	Date int `json:"ds"` // 日期，格式YYYYMMDD，不得早于 20240101
}

// GetLiveListResponse 获取直播大屏直播列表响应
type GetLiveListResponse struct {
	wxm.Error
	TraceId   string      `json:"trace_id,omitempty"`   // 追踪ID，报bug带
	LiveItems []*LiveItem `json:"live_items,omitempty"` // 直播信息
	HasMore   bool        `json:"has_more"`             // 是否还有更多的直播
}

// LiveItem 直播信息
type LiveItem struct {
	ExportId   string `json:"export_id"`   // 直播唯一ID
	CreateTime int64  `json:"create_time"` // 直播创建时间
}
