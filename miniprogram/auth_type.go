package miniprogram

import (
	"github.com/smartwalle/wxm"
)

type JSCode2SessionResponse struct {
	wxm.Error
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
}
