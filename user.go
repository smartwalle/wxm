package wxm

import (
	"context"
	"net/url"
)

const (
	kGetUserBaseInfo = "https://api.weixin.qq.com/sns/userinfo"
)

// GetUserBaseInfo 获取用户信息
//
//	公众号 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
//	微信 https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (c *Client) GetUserBaseInfo(ctx context.Context, accessToken, openId string, lang string) (result *GetUserBaseInfoResponse, err error) {
	var query = url.Values{}
	query.Add("access_token", accessToken)
	query.Add("openid", openId)
	query.Add("lang", lang)

	if err = c.Get(ctx, "", kGetUserBaseInfo, query, &result); err != nil {
		return nil, err
	}
	return result, nil
}
