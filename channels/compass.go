package channels

import "context"

const (
	APIGetFinderOverall = "/channels/ec/compass/finder/overall/get"
)

// GetFinderOverall 获取电商概览数据
//
//	接口文档：https://developers.weixin.qq.com/doc/channels/api/channels/compass/api_getfinderoverall.html
func (c *Channels) GetFinderOverall(ctx context.Context, accessToken string, request GetFinderOverallRequest) (response *GetFinderOverallResponse, err error) {
	if err = c.Post(ctx, APIGetFinderOverall, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
