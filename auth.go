package wxm

import (
	"context"
	"net/url"
)

const (
	kAccessToken  = "https://api.weixin.qq.com/sns/oauth2/access_token"
	kRefreshToken = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
)

// GetAccessToken 通过 Code 获取 AccessToken
//
//	公众号 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
//	网站 https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
//	微信 https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (c *Client) GetAccessToken(ctx context.Context, appId, secret, code string) (token *AccessToken, err error) {
	var query = url.Values{}
	query.Add("appid", appId)
	query.Add("secret", secret)
	query.Add("code", code)
	query.Add("grant_type", "authorization_code")

	if err = c.Get(ctx, "", kAccessToken, query, &token); err != nil {
		return nil, err
	}
	return token, nil
}

// RefreshAccessToken 刷新 AccessToken
//
//	公众号 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
//	网站 https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
//	微信 https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (c *Client) RefreshAccessToken(ctx context.Context, appId, refreshToken string) (token *RefreshToken, err error) {
	var query = url.Values{}
	query.Add("appid", appId)
	query.Add("refresh_token", refreshToken)
	query.Add("grant_type", "refresh_token")

	if err = c.Get(ctx, "", kRefreshToken, query, &token); err != nil {
		return nil, err
	}
	return token, nil
}
