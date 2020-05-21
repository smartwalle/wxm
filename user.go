package wxm

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	kGetUserOpenIdListURL = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s"
	kGetUserInfoURL       = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s"
)

// GetUserOpenIdList 公众号-获取帐号的关注者列表 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (this *Client) GetUserOpenIdList(nextOpenId string) (result *GetUserOpenIdListRsp, err error) {
	var values = url.Values{}
	values.Add("next_openid", nextOpenId)

	data, err := this.request(http.MethodGet, kGetUserOpenIdListURL, nil, values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserInfo 公众号-获取用户基本信息 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (this *Client) GetUserInfo(openId, lang string) (result *GetUserInfoRsp, err error) {
	var values = url.Values{}
	values.Add("openid", openId)
	values.Add("lang", lang)

	data, err := this.request(http.MethodGet, kGetUserInfoURL, nil, values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
