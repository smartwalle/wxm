package officialaccount

import "github.com/smartwalle/wxm"

type GetUserOpenIdListResponse struct {
	wxm.Error
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenId []string `json:"openid"`
	} `json:"data"`
	NextOpenId string `json:"next_openid"`
}

type UserInfo struct {
	Subscribe int    `json:"subscribe"`
	OpenId    string `json:"openid"`
	Nickname  string `json:"nickname"`
	//Sex            int    `json:"sex"`
	Language string `json:"language"`
	//City           string `json:"city"`
	//Province       string `json:"province"`
	//Country        string `json:"country"`
	HeadImgURL     string  `json:"headimgurl"`
	SubscribeTime  int64   `json:"subscribe_time"`
	UnionId        string  `json:"unionid"` // 只有在将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string  `json:"remark"`
	GroupId        int64   `json:"groupid"`
	TagIdList      []int64 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QRScene        int64   `json:"qr_scene"`
	QRSceneStr     string  `json:"qr_scene_str"`
}

type GetUserInfoResponse struct {
	wxm.Error
	UserInfo
}

type GetUserInfoListResponse struct {
	wxm.Error
	UserInfoList []*UserInfo `json:"user_info_list"`
}
