package miniprogram

import "github.com/smartwalle/wxm"

// GetUserPhoneNumberResponse 获取手机号返回数据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
type GetUserPhoneNumberResponse struct {
	wxm.Error
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

type UserInfoResponse struct {
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
