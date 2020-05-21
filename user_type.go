package wxm

type GetUserOpenIdListRsp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Data    struct {
		OpenId []string `json:"openid"`
	} `json:"data"`
	NextOpenId string `json:"next_openid"`
}

type GetUserInfoRsp struct {
	ErrCode        ErrCode `json:"errcode"`
	ErrMsg         string  `json:"errmsg"`
	Subscribe      int     `json"subscribe"`
	OpenId         string  `json"openid"`
	Nickname       string  `json"nickname"`
	Sex            int     `json"sex"`
	Language       string  `json"language"`
	City           string  `json"city"`
	Province       string  `json"province"`
	Country        string  `json"country"`
	HeadImgURL     string  `json"headimgurl"`
	SubscribeTime  int64   `json"subscribe_time"`
	UnionId        string  `json"unionid"`
	Remark         string  `json"remark"`
	GroupId        int64   `json"groupid"`
	TagIdList      int64   `json"tagid_list"`
	SubscribeScene string  `json"subscribe_scene"`
	QRScene        int64   `json"qr_scene"`
	QRSceneStr     string  `json"qr_scene_str"`
}
