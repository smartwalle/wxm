package channel

import "github.com/smartwalle/wxm"

// GetFinderAttrResponse 通过 AppID 获取视频号属性响应
type GetFinderAttrResponse struct {
	wxm.Error
	FinderAttr *FinderAttr `json:"finder_attr"`
}

// FinderAttr 视频号属性
type FinderAttr struct {
	UniqId    string `json:"uniq_id"`
	Nickname  string `json:"nickname"`
	FansCount int    `json:"fans_count"`
}

// GetFinderLiveRecordListResponse 获取视频号直播记录列表响应
type GetFinderLiveRecordListResponse struct {
	wxm.Error
	LiveRecordList []*LiveRecord `json:"live_record_list"`
	TotalCount     int           `json:"total_count"`
}

// LiveRecord 直播记录
type LiveRecord struct {
	ExportId   string `json:"export_id"`   // 直播 ID
	CreateTime int64  `json:"create_time"` // 创建时间
	ExpireTime int64  `json:"expire_time"` // 过期时间
	LiveStatus int    `json:"live_status"` // 直播状态：0-未开始，1-直播中，2-已结束，3-已过期，4-已取消
	Name       string `json:"name"`        // 直播标题
	CoverImg   string `json:"cover_img"`   // 直播封面
}
