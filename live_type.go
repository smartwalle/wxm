package wxm

type GetLiveInfoParam struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
}

type GetLiveInfoRsp struct {
	ErrCode  int             `json:"errcode"`
	ErrMsg   string          `json:"errmsg"`
	RoomInfo []*LiveRoomInfo `json:"room_info"`
	Total    int             `json:"total"`
}

type LiveRoomInfo struct {
	Name       string           `json:"name"`
	RoomId     int64            `json:"room_id"`
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
