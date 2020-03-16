package wxm

import (
	"fmt"
	"time"
)

type ErrCode int

const (
	CodeSuccess           ErrCode = 0     // 请求成功
	CodeInvalidCredential ErrCode = 40001 // access_token 无效或者 AppSecret 错误
	CodeInvalidGrantType  ErrCode = 40002 // 请确保 grant_type 字段值为 client_credential
	CodeInvalidAppId      ErrCode = 40013 // 不合法的 AppID，请开发者检查 AppID 的正确性，避免异常字符，注意大小写
	CodeInvalidCode       ErrCode = 40029 // 不合法的 code
	CodeMaxRate           ErrCode = 45009 // 调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成。
	CodeInvalidPage       ErrCode = 41030 // 所传page页面不存在，或者小程序没有发布
)

type Error struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
}

func (this *Error) Error() string {
	return fmt.Sprintf("%d-%s", this.ErrCode, this.ErrMsg)
}

type AccessToken struct {
	ErrCode     ErrCode `json:"errcode"`
	ErrMsg      string  `json:"errmsg"`
	AccessToken string  `json:"access_token"`
	ExpiresIn   int64   `json:"expires_in"`
	CreateTime  int64   `json:"create_time"`
}

func (this *AccessToken) Valid() bool {
	var now = time.Now().Unix()
	if now < this.CreateTime+this.ExpiresIn-10 {
		return true
	}
	return false
}

type JSCode2SessionRsp struct {
	ErrCode    ErrCode `json:"errcode"`
	ErrMsg     string  `json:"errmsg"`
	SessionKey string  `json:"session_key"`
	OpenId     int     `json:"open_id"`
}

type GetUnlimitedParam struct {
	Scene     string     `json:"scene"`
	Page      string     `json:"page"`
	Width     int        `json:"width,omitempty"`
	AutoColor bool       `json:"auto_color"`
	LineColor *LineColor `json:"line_color,omitempty"`
	IsHyaline bool       `json:"is_hyaline"`
}

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type GetUnlimitedRsp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
	Data    []byte  `json:"data"`
}
