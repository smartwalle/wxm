package wxm

// GetUserPhoneNumberResponse 获取手机号返回数据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
type GetUserPhoneNumberResponse struct {
	Error
	PhoneInfo *PhoneInfo `json:"phone_info"` // 用户手机号信息
}

type PhoneInfo struct {
	PhoneNumber     string     `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string     `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string     `json:"countryCode"`     // 区号
	Watermark       *Watermark `json:"watermark"`       // 数据水印
}

type Watermark struct {
	AppId     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type MiniProgramUserInfoResponse struct {
	OpenId    string     `json:"openid"`
	Nickname  string     `json:"nickname"`
	Gender    int        `json:"gender"`
	Language  string     `json:"language"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	Country   string     `json:"country"`
	AvatarURL string     `json:"avatarUrl"`
	UnionId   string     `json:"unionid"`
	Watermark *Watermark `json:"watermark"`
}

type GetUserOpenIdListResponse struct {
	Error
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenId []string `json:"openid"`
	} `json:"data"`
	NextOpenId string `json:"next_openid"`
}

type GetUserBaseInfoResponse struct {
	Error
	UserBaseInfo
}

type UserBaseInfo struct {
	OpenId   string `json:"openid"`
	Nickname string `json:"nickname"`
	//Sex        int      `json:"sex"`
	//City       string   `json:"city"`
	//Province   string   `json:"province"`
	//Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	UnionId    string   `json:"unionid"` // 只有在将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Privilege  []string `json:"privilege"`
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
	Error
	UserInfo
}

type GetUserInfoListParam struct {
	UserList []map[string]string `json:"user_list"`
}

func (u *GetUserInfoListParam) AddOpenId(openIds ...string) {
	if len(openIds) == 0 {
		return
	}
	if len(u.UserList) == 0 {
		u.UserList = make([]map[string]string, 0, len(openIds))
	}

	for _, openId := range openIds {
		u.UserList = append(u.UserList, map[string]string{"openid": openId})
	}
}

type GetUserInfoListResponse struct {
	Error
	UserInfoList []*UserInfo `json:"user_info_list"`
}
