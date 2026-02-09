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
