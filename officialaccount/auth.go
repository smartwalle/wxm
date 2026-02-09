package officialaccount

import (
	"net/url"
)

const (
	APIAuthorize = "https://open.weixin.qq.com/connect/oauth2/authorize"
)

// GetAuthorizeURL 获取公众号登录 URL
//
//	接口文档：https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
//
//	1. 服务端调用 GetAuthorizeURL 生成登录 URL，微信 APP 或者浏览器 中访问该 URL 成功之后，会重定向到 redirectURL
//	2. 服务端对应的 redirectURL 接口获取 code 参数，然后调用 GetAccessToken 获取 AccessToken 信息
func (o *OfficialAccount) GetAuthorizeURL(appId, redirectURL string, scope AuthScope, state string) string {
	var query = url.Values{}
	query.Add("appid", appId)
	query.Add("redirect_uri", redirectURL)
	query.Add("response_type", "code")
	query.Add("scope", string(scope))
	query.Add("state", state)
	return APIAuthorize + "?" + query.Encode()
}
