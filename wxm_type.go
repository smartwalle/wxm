package wxm

import "time"

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (this *Error) Error() string {
	return this.ErrMsg
}

type AccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	CreateTime  int64  `json:"create_time"`
}

func (this *AccessToken) Valid() bool {
	var now = time.Now().Unix()
	if now < this.CreateTime+this.ExpiresIn-10 {
		return true
	}
	return false
}

type GetUnlimitedParam struct {
	Scene     string `json:"scene"`
	Page      string `json:"page"`
	Width     int    `json:"width,omitempty"`
	AutoColor bool   `json:"auto_color"`
	IsHyaline bool   `json:"is_hyaline"`
}
