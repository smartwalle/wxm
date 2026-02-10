package officialaccount

import (
	"context"
	"net/url"
)

const (
	APIGetUseList      = "/cgi-bin/user/get"
	APIGetUserInfo     = "/cgi-bin/user/info"
	APIGetUserInfoList = "/cgi-bin/user/info/batchget"
)

// GetUserList 获取帐号的关注者列表
//
//	接口文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (o *OfficialAccount) GetUserList(ctx context.Context, accessToken, nextOpenId string) (response *GetUserListResponse, err error) {
	var query = url.Values{}
	if nextOpenId != "" {
		query.Add("next_openid", nextOpenId)
	}

	if err = o.Get(ctx, APIGetUseList, accessToken, query, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetUserInfo 获取用户基本信息
//
//	接口文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (o *OfficialAccount) GetUserInfo(ctx context.Context, accessToken, openId, lang string) (response *GetUserInfoResponse, err error) {
	var query = url.Values{}
	query.Add("openid", openId)
	query.Add("lang", lang)

	if err = o.Get(ctx, APIGetUserInfo, accessToken, query, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetUserInfoList 批量获取用户基本信息
func (o *OfficialAccount) GetUserInfoList(ctx context.Context, accessToken string, openIds []string) (response *GetUserInfoListResponse, err error) {
	if len(openIds) == 0 {
		return &GetUserInfoListResponse{}, nil
	}

	var request = struct {
		UserList []map[string]string `json:"user_list"`
	}{
		UserList: make([]map[string]string, 0, len(openIds)),
	}
	for _, openId := range openIds {
		request.UserList = append(request.UserList, map[string]string{"openid": openId})
	}

	if err = o.Post(ctx, APIGetUserInfoList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
