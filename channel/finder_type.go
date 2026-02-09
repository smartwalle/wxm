package channel

import "github.com/smartwalle/wxm"

// GetFinderAttrResponse 获取视频号账号信息响应
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

// GetFinderLiveRecordListResponse 视频号获取当前的直播记录响应
type GetFinderLiveRecordListResponse struct {
	wxm.Error
	LiveList []*LiveRecord `json:"live_list"`
}

// LiveRecord 直播记录
type LiveRecord struct {
	ExportId    string `json:"export_id"`   // 直播id
	Description string `json:"description"` // 直播描述
	CoverUrl    string `json:"cover_url"`   // 直播封面
	Nickname    string `json:"nickname"`    // 开播视频号昵称
	HeadUrl     string `json:"head_url"`    // 开播视频号头像
}
