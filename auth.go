package wxm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	kAuthorizeURL      = "https://open.weixin.qq.com/connect/oauth2/authorize"
	kQRConnectURL      = "https://open.weixin.qq.com/connect/qrconnect"
	kAccessTokenURL    = "https://api.weixin.qq.com/sns/oauth2/access_token"
	kJSCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?grant_type=%s&appid=%s&secret=%s&js_code=%s"
)

// 1. 服务端调用 GetAuthorizeURL 或者 GetQRConnectURL 生成登录 URL，微信 APP 或者浏览器 中访问该 URL 成功之后，会重定向到 redirectURL
// 2. 服务端对应的 redirectURL 接口获取 code 参数，然后调用 GetAccessToken 获取 AccessToken 信息

// GetAuthorizeURL 微信用户-获取公众号登录 URL
func (this *OfficialAccount) GetAuthorizeURL(redirectURL string, scope AuthScope, state string) string {
	var v = url.Values{}
	v.Add("appid", this.client.appId)
	v.Add("redirect_uri", redirectURL)
	v.Add("response_type", "code")
	v.Add("scope", string(scope))
	v.Add("state", state)
	return kAuthorizeURL + "?" + v.Encode()
}

// GetQRConnectURL 网站-获取网站应用微信登录 URL https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func (this *Website) GetQRConnectURL(redirectURL string, state string) string {
	var v = url.Values{}
	v.Add("appid", this.client.appId)
	v.Add("redirect_uri", redirectURL)
	v.Add("response_type", "code")
	v.Add("scope", "snsapi_login")
	v.Add("state", state)
	return kQRConnectURL + "?" + v.Encode()
}

// GetAccessToken 微信用户-通过 Code 获取 AccessToken https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (this *client) GetAccessToken(code string) (result *AccessToken, err error) {
	var v = url.Values{}
	v.Add("appid", this.appId)
	v.Add("secret", this.appSecret)
	v.Add("secret", this.appSecret)
	v.Add("code", code)
	v.Add("grant_type", "authorization_code")

	var nURL = kAccessTokenURL + "?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, nURL, nil)
	if err != nil {
		return nil, err
	}
	rsp, err := this.client.Do(req)
	if rsp != nil && rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *OfficialAccount) GetAccessToken(code string) (result *AccessToken, err error) {
	return this.client.GetAccessToken(code)
}

func (this *Website) GetAccessToken(code string) (result *AccessToken, err error) {
	return this.client.GetAccessToken(code)
}

// JSCode2Session 小程序-登录 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (this *MiniProgram) JSCode2Session(code string) (result *JSCode2SessionRsp, err error) {
	var nURL = fmt.Sprintf(kJSCode2SessionURL, "authorization_code", this.client.appId, this.client.appSecret, code)

	req, err := http.NewRequest(http.MethodGet, nURL, nil)
	if err != nil {
		return nil, err
	}
	rsp, err := this.client.client.Do(req)
	if rsp != nil && rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
