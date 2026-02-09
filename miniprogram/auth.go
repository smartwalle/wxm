package miniprogram

import (
	"context"
	"net/url"
)

const APIJSCode2Session = "/sns/jscode2session"

// JSCode2Session 登录凭证校验
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (m *MiniProgram) JSCode2Session(ctx context.Context, appId, secret, code string) (response *JSCode2SessionResponse, err error) {
	var query = url.Values{}
	query.Add("appid", appId)
	query.Add("secret", secret)
	query.Add("js_code", code)
	query.Add("grant_type", "authorization_code")

	if err = m.Get(ctx, APIJSCode2Session, "", query, &response); err != nil {
		return nil, err
	}
	return response, nil
}
