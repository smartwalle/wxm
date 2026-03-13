package channels

import "context"

const (
	APIGetLiveList = "/channels/livedashboard/getlivelist"
)

// GetLiveList 获取直播大屏直播列表
//
//	接口文档：https://developers.weixin.qq.com/doc/channels/api/channels/livedashboard/api_getlivelist.html
func (c *Channels) GetLiveList(ctx context.Context, accessToken string, request GetLiveListRequest) (response *GetLiveListResponse, err error) {
	if err = c.Post(ctx, APIGetLiveList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
