package wxm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	kGetUserPhoneNumber = "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
	kGetUserBaseInfo    = "https://api.weixin.qq.com/sns/userinfo"
	kGetUserOpenIdList  = "https://api.weixin.qq.com/cgi-bin/user/get"
	kGetUserInfo        = "https://api.weixin.qq.com/cgi-bin/user/info"
	kGetUserInfoList    = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
)

// GetUserPhoneNumber 获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func (m *MiniProgram) GetUserPhoneNumber(ctx context.Context, code string) (result *GetUserPhoneNumberResponse, err error) {
	var param = struct {
		Code string `json:"code"`
	}{
		Code: code,
	}
	if err = m.client.requestWithAccessToken(ctx, http.MethodPost, kGetUserPhoneNumber, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DecodePhoneNumber 小程序-解密手机号码数据 https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/deprecatedGetPhoneNumber.html
//
// 小程序端申请获取用户的手机号码之后，获取到的是加密的数据，需要调用本方法对该数据进行解密，以获取手机号码。
func (m *MiniProgram) DecodePhoneNumber(sessionKey, encryptedData, iv string) (result *PhoneInfo, err error) {
	plaintext, err := m.decrypt(sessionKey, encryptedData, iv)
	if err = json.Unmarshal(plaintext, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DecodeUserInfo 小程序-解密用户数据 https://developers.weixin.qq.com/miniprogram/dev/api/open-api/user-info/wx.getUserInfo.html
//
// 小程序端申请获取用户的信息之后，获取到的有加密的数据，需要调用本方法对该数据进行解密，以获取加密信息。
func (m *MiniProgram) DecodeUserInfo(sessionKey, encryptedData, iv string) (result *MiniProgramUserInfoResponse, err error) {
	plaintext, err := m.decrypt(sessionKey, encryptedData, iv)
	if err = json.Unmarshal(plaintext, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetUserBaseInfo(ctx context.Context, accessToken, openId string, lang string) (result *GetUserBaseInfoResponse, err error) {
	var v = url.Values{}
	v.Add("access_token", accessToken)
	v.Add("openid", openId)
	v.Add("lang", lang)

	if err = c.requestWithoutAccessToken(ctx, http.MethodGet, kGetUserBaseInfo, nil, v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserBaseInfo 公众号-获取用户信息 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func (o *OfficialAccount) GetUserBaseInfo(ctx context.Context, accessToken, openId, lang string) (result *GetUserBaseInfoResponse, err error) {
	return o.client.GetUserBaseInfo(ctx, accessToken, openId, lang)
}

// GetUserBaseInfo 微信-获取用户信息 https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (m *MobileApp) GetUserBaseInfo(ctx context.Context, accessToken, openId, lang string) (result *GetUserBaseInfoResponse, err error) {
	return m.client.GetUserBaseInfo(ctx, accessToken, openId, lang)
}

// GetUserOpenIdList 公众号-获取帐号的关注者列表 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (o *OfficialAccount) GetUserOpenIdList(ctx context.Context, nextOpenId string) (result *GetUserOpenIdListResponse, err error) {
	var v = url.Values{}
	v.Add("next_openid", nextOpenId)

	if err = o.client.requestWithAccessToken(ctx, http.MethodGet, kGetUserOpenIdList, nil, v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserInfo 公众号-获取用户基本信息 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (o *OfficialAccount) GetUserInfo(ctx context.Context, openId, lang string) (result *GetUserInfoResponse, err error) {
	var v = url.Values{}
	v.Add("openid", openId)
	v.Add("lang", lang)

	if err = o.client.requestWithAccessToken(ctx, http.MethodGet, kGetUserInfo, nil, v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserInfoList 公众号-批量获取用户基本信息
func (o *OfficialAccount) GetUserInfoList(ctx context.Context, openIds ...string) (result *GetUserInfoListResponse, err error) {
	if len(openIds) == 0 {
		return &GetUserInfoListResponse{}, nil
	}

	var param = struct {
		UserList []map[string]string `json:"user_list"`
	}{
		UserList: make([]map[string]string, 0, len(openIds)),
	}
	for _, openId := range openIds {
		param.UserList = append(param.UserList, map[string]string{"openid": openId})
	}

	if err = o.client.requestWithAccessToken(ctx, http.MethodPost, kGetUserInfoList, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
