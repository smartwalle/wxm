package wxm

type Watermark struct {
	AppId     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type GetPhoneNumberRsp struct {
	PhoneNumber     string     `json:"phoneNumber"`
	PurePhoneNumber string     `json:"purePhoneNumber"`
	CountryCode     string     `json:"countryCode"`
	Watermark       *Watermark `json:"watermark"`
}

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

type GetUserBaseInfoRsp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
	UserBaseInfo
}

type UserBaseInfo struct {
	OpenId     string   `json"openid"`
	Nickname   string   `json"nickname"`
	Sex        int      `json"sex"`
	City       string   `json"city"`
	Province   string   `json"province"`
	Country    string   `json"country"`
	HeadImgURL string   `json"headimgurl"`
	UnionId    string   `json"unionid"`
	Privilege  []string `json:"privilege"`
}

type UserInfo struct {
	Subscribe      int    `json"subscribe"`
	OpenId         string `json"openid"`
	Nickname       string `json"nickname"`
	Sex            int    `json"sex"`
	Language       string `json"language"`
	City           string `json"city"`
	Province       string `json"province"`
	Country        string `json"country"`
	HeadImgURL     string `json"headimgurl"`
	SubscribeTime  int64  `json"subscribe_time"`
	UnionId        string `json"unionid"`
	Remark         string `json"remark"`
	GroupId        int64  `json"groupid"`
	TagIdList      int64  `json"tagid_list"`
	SubscribeScene string `json"subscribe_scene"`
	QRScene        int64  `json"qr_scene"`
	QRSceneStr     string `json"qr_scene_str"`
}

type GetUserInfoRsp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
	UserInfo
}

type GetUserInfoListParam struct {
	UserList []map[string]string `json:"user_list"`
}

func (this *GetUserInfoListParam) AddOpenId(openIds ...string) {
	if len(openIds) == 0 {
		return
	}
	if len(this.UserList) == 0 {
		this.UserList = make([]map[string]string, 0, len(openIds))
	}

	for _, openId := range openIds {
		this.UserList = append(this.UserList, map[string]string{"openid": openId})
	}
}

type GetUserInfoListRsp struct {
	ErrCode      ErrCode     `json:"errcode"`
	ErrMsg       string      `json:"errmsg"`
	UserInfoList []*UserInfo `json:"user_info_list"`
}
