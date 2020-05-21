package wxm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	kAuthorizeURL   = "https://open.weixin.qq.com/connect/oauth2/authorize"
	kAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"
)

// 1. 服务端调用 GetAuthorizeURL 生成登录 URL，微信 APP 中访问该 URL 成功之后，会重定向到 redirectURL
// 2. 服务端对应的 redirectURL 接口获取 code 参数，然后调用 GetAccessToken 获取 AccessToken 信息

// GetAuthorizeURL 获取公众号登录 URL
func (this *Client) GetAuthorizeURL(redirectURL string, scope AuthScope, state string) string {
	var v = url.Values{}
	v.Add("appid", this.appId)
	v.Add("redirect_uri", redirectURL)
	v.Add("response_type", "code")
	v.Add("scope", string(scope))
	v.Add("state", state)
	return kAuthorizeURL + "?" + v.Encode()
}

// GetAccessToken 获取 AccessToken
func (this *Client) GetAccessToken(code string) (result *AccessToken, err error) {
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
