package store

import "github.com/smartwalle/wxm"

// BasicInfo 店铺基本信息
type BasicInfo struct {
	Nickname      string `json:"nickname"`       // 店铺名称
	HeadImgURL    string `json:"headimg_url"`    // 店铺头像URL
	SubjectType   string `json:"subject_type"`   // 店铺类型，目前为"企业"或"个体工商户"
	Status        string `json:"status"`         // 店铺状态，目前为"opening"或"open_finished"或"closing"或"close_finished"
	Username      string `json:"username"`       // 店铺原始ID
	IsLocalLife   bool   `json:"is_local_life"`  // 是否本地生活小店
	OpenTimestamp int64  `json:"open_timestamp"` // 店铺状态为"open_finished"时，返回开店时间戳
}

// GetBasicInfoResponse 获取店铺基本信息响应
type GetBasicInfoResponse struct {
	wxm.Error
	Info *BasicInfo `json:"info"` // 店铺信息
}
