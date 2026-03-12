package channels

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

// GetFinderLiveDataListRequest 获取留资直播数据详情请求参数
type GetFinderLiveDataListRequest struct {
	StartTime  int64  `json:"start_time,omitempty"`  // 查询范围的开始时间（与关播时间对比），时间戳
	EndTime    int64  `json:"end_time,omitempty"`    // 查询范围的结束时间（与关播时间对比），时间戳
	LastBuffer string `json:"last_buffer,omitempty"` // 顺序翻页，传入上次请求返回的last_buffer，会从上次返回的结果往后翻一页
}

// GetFinderLiveDataListResponse 获取留资直播数据详情响应
type GetFinderLiveDataListResponse struct {
	wxm.Error
	LastBuffer   string            `json:"last_buffer,omitempty"`   // 本次翻页的上下文，用于顺序翻页请求
	ContinueFlag bool              `json:"continue_flag,omitempty"` // 是否还有直播
	Item         []*FinderLiveData `json:"item,omitempty"`          // 直播信息
}

// FinderLiveData 留资直播数据
type FinderLiveData struct {
	// 基础信息
	ExportId            string `json:"export_id,omitempty"`                // 直播唯一id
	LiveStartTime       int64  `json:"live_start_time,omitempty"`          // 开播时间戳（秒）
	LiveDurationSeconds int64  `json:"live_duration_in_seconds,omitempty"` // 直播时长（秒）

	// 统计数据
	TotalAudienceCount  int64 `json:"total_audience_count,omitempty"`  // 观看人数
	TotalCheerCount     int64 `json:"total_cheer_count,omitempty"`     // 喝彩次数（点赞数）
	ForwardCount        int64 `json:"forward_count,omitempty"`         // 分享次数
	TotalCommentCount   int64 `json:"total_comment_count,omitempty"`   // 评论条数
	AudiencesAvgSeconds int64 `json:"audiences_avg_seconds,omitempty"` // 人均观看时长（秒）
	MaxOnlineCount      int64 `json:"max_online_count,omitempty"`      // 最高在线人数
	NewFollowCount      int64 `json:"new_follow_count,omitempty"`      // 新增粉丝
	NewFollowCountBiz   int64 `json:"new_follow_count_biz,omitempty"`  // 公众号新增粉丝
}
