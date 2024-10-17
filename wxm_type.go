package wxm

import (
	"fmt"
	"time"
)

type Code int

func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}

func (c Code) IsFailure() bool {
	return c != CodeSuccess
}

const (
	CodeSuccess           Code = 0     // 请求成功
	CodeInvalidCredential Code = 40001 // access_token 无效或者 AppSecret 错误
	CodeInvalidGrantType  Code = 40002 // 请确保 grant_type 字段值为 client_credential
	CodeInvalidAppId      Code = 40013 // 不合法的 AppID，请开发者检查 AppID 的正确性，避免异常字符，注意大小写
	CodeInvalidCode       Code = 40029 // 不合法的 code
	CodeMaxRate           Code = 45009 // 调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成。
	CodeInvalidPage       Code = 41030 // 所传page页面不存在，或者小程序没有发布
)

type Error struct {
	Code Code   `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d-%s", e.Code, e.Msg)
}

func (e Error) IsSuccess() bool {
	return e.Code.IsSuccess()
}

func (e Error) IsFailure() bool {
	return e.Code.IsFailure()
}

type Token struct {
	Error
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	CreateTime  int64  `json:"create_time"`
}

func (t *Token) Valid() bool {
	var now = time.Now().Unix()
	if now < t.CreateTime+t.ExpiresIn-60 {
		return true
	}
	return false
}
