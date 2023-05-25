package wxm

import (
	"net/http"
	"net/url"
)

const (
	kAuthorize      = "https://open.weixin.qq.com/connect/oauth2/authorize"
	kQRConnect      = "https://open.weixin.qq.com/connect/qrconnect"
	kAccessToken    = "https://api.weixin.qq.com/sns/oauth2/access_token"
	kRefreshToken   = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	kJSCode2Session = "https://api.weixin.qq.com/sns/jscode2session"
)

// 1. 服务端调用 GetAuthorizeURL 或者 GetQRConnectURL 生成登录 URL，微信 APP 或者浏览器 中访问该 URL 成功之后，会重定向到 redirectURL
// 2. 服务端对应的 redirectURL 接口获取 code 参数，然后调用 GetAccessToken 获取 AccessToken 信息

// GetAuthorizeURL 公众号-获取公众号登录 URL https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func (this *OfficialAccount) GetAuthorizeURL(redirectURL string, scope AuthScope, state string) string {
	var v = url.Values{}
	v.Add("appid", this.client.appId)
	v.Add("redirect_uri", redirectURL)
	v.Add("response_type", "code")
	v.Add("scope", string(scope))
	v.Add("state", state)
	return kAuthorize + "?" + v.Encode()
}

// GetQRConnectURL 网站-获取网站应用微信登录 URL https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func (this *Website) GetQRConnectURL(redirectURL string, state string) string {
	var v = url.Values{}
	v.Add("appid", this.client.appId)
	v.Add("redirect_uri", redirectURL)
	v.Add("response_type", "code")
	v.Add("scope", "snsapi_login")
	v.Add("state", state)
	return kQRConnect + "?" + v.Encode()
}

// GetAccessToken 通过 Code 获取 AccessToken
func (this *client) GetAccessToken(code string) (result *AccessToken, err error) {
	var v = url.Values{}
	v.Add("appid", this.appId)
	v.Add("secret", this.appSecret)
	v.Add("code", code)
	v.Add("grant_type", "authorization_code")

	if err = this.requestWithoutAccessToken(http.MethodGet, kAccessToken, nil, v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccessToken 公众号-获取 AccessToken https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func (this *OfficialAccount) GetAccessToken(code string) (result *AccessToken, err error) {
	return this.client.GetAccessToken(code)
}

// GetAccessToken 网站-获取 AccessToken https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func (this *Website) GetAccessToken(code string) (result *AccessToken, err error) {
	return this.client.GetAccessToken(code)
}

// GetAccessToken 微信-获取 AccessToken https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (this *MobileApp) GetAccessToken(code string) (result *AccessToken, err error) {
	return this.client.GetAccessToken(code)
}

func (this *client) RefreshAccessToken(refreshToken string) (result *RefreshToken, err error) {
	var v = url.Values{}
	v.Add("appid", this.appId)
	v.Add("refresh_token", refreshToken)
	v.Add("grant_type", "refresh_token")

	if err = this.requestWithoutAccessToken(http.MethodGet, kRefreshToken, nil, v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// RefreshAccessToken 公众号-刷新 AccessToken https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func (this *OfficialAccount) RefreshAccessToken(refreshToken string) (result *RefreshToken, err error) {
	return this.client.RefreshAccessToken(refreshToken)
}

// RefreshAccessToken 网站-刷新 AccessToken https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func (this *Website) RefreshAccessToken(refreshToken string) (result *RefreshToken, err error) {
	return this.client.RefreshAccessToken(refreshToken)
}

// RefreshAccessToken 微信-刷新 AccessToken https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (this *MobileApp) RefreshAccessToken(refreshToken string) (result *RefreshToken, err error) {
	return this.client.RefreshAccessToken(refreshToken)
}

// JSCode2Session 小程序-登录凭证校验 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (this *MiniProgram) JSCode2Session(code string) (result *JSCode2SessionRsp, err error) {
	var v = url.Values{}
	v.Add("appid", this.client.appId)
	v.Add("secret", this.client.appSecret)
	v.Add("js_code", code)
	v.Add("grant_type", "authorization_code")

	if err = this.client.requestWithoutAccessToken(http.MethodGet, kJSCode2Session, nil, v, &result); err != nil {
		return nil, err
	}
	return result, nil
}
