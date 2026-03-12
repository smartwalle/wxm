package channels

import (
	"context"
)

const (
	APIGetFinderAttr           = "/channels/finderlive/get_finder_attr_by_appid"
	APIGetFinderLiveRecordList = "/channels/ec/finderlive/getfinderliverecordlist"
	APIGetFinderLiveDataList   = "/channels/ec/finderlive/get_finder_live_data_list"
)

// GetFinderAttr 获取视频号账号信息
//
//	接口文档：https://developers.weixin.qq.com/doc/channels/api/channels/leadslive/api_getfinderattrbyappid.html
func (c *Channels) GetFinderAttr(ctx context.Context, accessToken string) (response *GetFinderAttrResponse, err error) {
	var request = struct{}{}
	if err = c.Post(ctx, APIGetFinderAttr, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetFinderLiveRecordList 视频号获取当前的直播记录
//
//	接口文档：https://developers.weixin.qq.com/doc/channels/api/channels/live/api_getfinderliverecordlist.html
func (c *Channels) GetFinderLiveRecordList(ctx context.Context, accessToken string) (response *GetFinderLiveRecordListResponse, err error) {
	var request = struct{}{}
	if err = c.Post(ctx, APIGetFinderLiveRecordList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetFinderLiveDataList 获取留资直播数据详情
//
//	接口文档：https://developers.weixin.qq.com/doc/channels/api/channels/leadslive/api_getfinderlivedatalist.html
func (c *Channels) GetFinderLiveDataList(ctx context.Context, accessToken string, request GetFinderLiveDataListRequest) (response *GetFinderLiveDataListResponse, err error) {
	if err = c.Post(ctx, APIGetFinderLiveDataList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
