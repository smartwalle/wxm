package wxm

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	kGetUserBaseInfoURL   = "https://api.weixin.qq.com/sns/userinfo"
	kGetUserOpenIdListURL = "https://api.weixin.qq.com/cgi-bin/user/get"
	kGetUserInfoURL       = "https://api.weixin.qq.com/cgi-bin/user/info"
	kGetUserInfoListURL   = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
)

func (this *client) GetUserBaseInfo(accessToken, openId string) (result *GetUserBaseInfoRsp, err error) {
	var values = url.Values{}
	values.Add("access_token", accessToken)
	values.Add("openid", openId)

	data, err := this.RequestWithoutAccessToken(http.MethodGet, kGetUserBaseInfoURL, nil, values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserBaseInfo 公众号-获取用户信息 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func (this *OfficialAccount) GetUserBaseInfo(accessToken, openId string) (result *GetUserBaseInfoRsp, err error) {
	return this.client.GetUserBaseInfo(accessToken, openId)
}

// GetUserBaseInfo 微信-获取用户信息 https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (this *MobileApp) GetUserBaseInfo(accessToken, openId string) (result *GetUserBaseInfoRsp, err error) {
	return this.client.GetUserBaseInfo(accessToken, openId)
}

// GetUserOpenIdList 公众号-获取帐号的关注者列表 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (this *OfficialAccount) GetUserOpenIdList(nextOpenId string) (result *GetUserOpenIdListRsp, err error) {
	var values = url.Values{}
	values.Add("next_openid", nextOpenId)

	data, err := this.client.RequestWithAccessToken(http.MethodGet, kGetUserOpenIdListURL, nil, values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserInfo 公众号-获取用户基本信息 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (this *OfficialAccount) GetUserInfo(openId, lang string) (result *GetUserInfoRsp, err error) {
	var values = url.Values{}
	values.Add("openid", openId)
	values.Add("lang", lang)

	data, err := this.client.RequestWithAccessToken(http.MethodGet, kGetUserInfoURL, nil, values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserInfoList 公众号-批量获取用户基本信息
func (this *OfficialAccount) GetUserInfoList(openIds ...string) (result *GetUserInfoListRsp, err error) {
	if len(openIds) == 0 {
		return &GetUserInfoListRsp{}, nil
	}

	var param = &GetUserInfoListParam{}
	param.AddOpenId(openIds...)

	data, err := this.client.RequestWithAccessToken(http.MethodPost, kGetUserInfoListURL, param, nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
