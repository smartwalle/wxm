package wxm

type OfficialAccount struct {
	*client
}

func NewOfficialAccount(appId, appSecret string) *OfficialAccount {
	var c = &OfficialAccount{}
	c.client = newClient(appId, appSecret)
	return c
}

// GetToken 公众号-获取全局唯一后台接口调用凭据（access_token）https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func (this *OfficialAccount) GetToken() (result string, err error) {
	return this.client.GetToken()
}

// RefreshToken 公众号-刷新本地全局唯一后台接口调用凭据（access_token）
func (this *OfficialAccount) RefreshToken() (err error) {
	return this.client.RefreshToken()
}
