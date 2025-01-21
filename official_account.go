package wxm

import (
	"context"
)

type OfficialAccount struct {
	client *client
}

func NewOfficialAccount(appId, appSecret string, opts ...Option) *OfficialAccount {
	var c = &OfficialAccount{}
	c.client = newClient(appId, appSecret, opts...)
	return c
}

func (o *OfficialAccount) With(opts ...Option) *OfficialAccount {
	var n = &OfficialAccount{}
	n.client = o.client.With(opts...)
	return n
}

func (o *OfficialAccount) SetAccessToken(accessToken string) {
	o.client.accessToken = accessToken
}

// GetToken 公众号-获取全局唯一后台接口调用凭据（access_token）https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func (o *OfficialAccount) GetToken(ctx context.Context) (token *Token, err error) {
	return o.client.GetToken(ctx)
}
