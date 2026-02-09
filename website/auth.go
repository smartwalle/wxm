package website

import (
	"net/url"
)

const (
	kQRConnect = "https://open.weixin.qq.com/connect/qrconnect"
)

// GetQRConnectURL 获取网站应用微信登录 URL
//
//	接口文档：https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
//
//	1. 服务端调用 GetQRConnectURL 生成登录 URL，微信 APP 或者浏览器 中访问该 URL 成功之后，会重定向到 redirectURL
//	2. 服务端对应的 redirectURL 接口获取 code 参数，然后调用 GetAccessToken 获取 AccessToken 信息
func (w *Website) GetQRConnectURL(appId, redirectURL string, state string) string {
	var query = url.Values{}
	query.Add("appid", appId)
	query.Add("redirect_uri", redirectURL)
	query.Add("response_type", "code")
	query.Add("scope", "snsapi_login")
	query.Add("state", state)
	return kQRConnect + "?" + query.Encode()
}
