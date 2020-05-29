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

// GetPhoneNumber 小程序-解密手机号码数据 https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html
//
// 小程序端申请获取用户的手机号码之后，获取到的是加密的数据，需要调用本方法对该数据进行解密，以获取手机号码。
func (this *MiniProgram) GetPhoneNumber(sessionKey, encryptedData, iv string) (result *MiniProgramPhoneNumber, err error) {
	decryptedBytes, err := this.decrypt(sessionKey, encryptedData, iv)
	if err = json.Unmarshal(decryptedBytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserInfo 小程序-解密用户数据 https://developers.weixin.qq.com/miniprogram/dev/api/open-api/user-info/wx.getUserInfo.html
//
// 小程序端申请获取用户的信息之后，获取到的有加密的数据，需要调用本方法对该数据进行解密，以获取加密信息。
func (this *MiniProgram) GetUserInfo(sessionKey, encryptedData, iv string) (result *MiniProgramUserInfo, err error) {
	decryptedBytes, err := this.decrypt(sessionKey, encryptedData, iv)
	if err = json.Unmarshal(decryptedBytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *client) GetUserBaseInfo(accessToken, openId string, lang string) (result *GetUserBaseInfoRsp, err error) {
	var v = url.Values{}
	v.Add("access_token", accessToken)
	v.Add("openid", openId)
	v.Add("lang", lang)

	data, err := this.RequestWithoutAccessToken(http.MethodGet, kGetUserBaseInfoURL, nil, v)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserBaseInfo 公众号-获取用户信息 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func (this *OfficialAccount) GetUserBaseInfo(accessToken, openId, lang string) (result *GetUserBaseInfoRsp, err error) {
	return this.client.GetUserBaseInfo(accessToken, openId, lang)
}

// GetUserBaseInfo 微信-获取用户信息 https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (this *MobileApp) GetUserBaseInfo(accessToken, openId, lang string) (result *GetUserBaseInfoRsp, err error) {
	return this.client.GetUserBaseInfo(accessToken, openId, lang)
}

// GetUserOpenIdList 公众号-获取帐号的关注者列表 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (this *OfficialAccount) GetUserOpenIdList(nextOpenId string) (result *GetUserOpenIdListRsp, err error) {
	var v = url.Values{}
	v.Add("next_openid", nextOpenId)

	data, err := this.client.RequestWithAccessToken(http.MethodGet, kGetUserOpenIdListURL, nil, v)
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
	var v = url.Values{}
	v.Add("openid", openId)
	v.Add("lang", lang)

	data, err := this.client.RequestWithAccessToken(http.MethodGet, kGetUserInfoURL, nil, v)
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
