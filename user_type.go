package wxm

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
