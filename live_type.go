package wxm

const (
	CodeEmptyRoomList Code = 1 // 没有房间信息
)

// GetLiveInfo https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/live-player-plugin.html
type GetLiveInfo struct {
	Action string `json:"action,omitempty"`  // 获取回放列表的时候需要传递字符串 get_replay
	RoomId int    `json:"room_id,omitempty"` // 获取回放列表的时候需要传递
	Start  int    `json:"start"`
	Limit  int    `json:"limit"`
}

type GetLiveInfoRsp struct {
	Error
	RoomInfo   []*LiveRoomInfo `json:"room_info"`
	LiveReplay []*LiveReplay   `json:"live_replay"`
	Total      int             `json:"total"`
}

type LiveRoomInfo struct {
	Name       string           `json:"name"`
	RoomId     int64            `json:"roomid"`
	CoverImg   string           `json:"cover_img"`
	LiveStatus int              `json:"live_status"`
	StartTime  int64            `json:"start_time"`
	EndTime    int64            `json:"end_time"`
	AnchorName string           `json:"anchor_name"`
	AnchorImg  string           `json:"anchor_img"`
	Goods      []*LiveGoodsInfo `json:"goods"`
}

type LiveGoodsInfo struct {
	CoverImg string `json:"cover_img"`
	URL      string `json:"url"`
	Price    int    `json:"price"`
	Name     string `json:"name"`
}

type LiveReplay struct {
	CreateTime string `json:"create_time"`
	ExpireTime string `json:"expire_time"`
	MediaURL   string `json:"media_url"`
}
