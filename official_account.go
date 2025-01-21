package wxm

import "net/http"

type OfficialAccount struct {
	client *client
}

func NewOfficialAccount(appId, appSecret string) *OfficialAccount {
	var c = &OfficialAccount{}
	c.client = newClient(appId, appSecret)
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

func (o *OfficialAccount) SetHTTPClient(client *http.Client) {
	o.client.client = client
}

// GetToken 公众号-获取全局唯一后台接口调用凭据（access_token）https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func (o *OfficialAccount) GetToken() (token *Token, err error) {
	return o.client.GetToken()
}
