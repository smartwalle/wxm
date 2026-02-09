package miniprogram

import (
	"context"
)

const (
	APIGetLiveInfo = "/wxa/business/getliveinfo"
)

// GetLiveInfo 获取直播房间列表
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/live-player-plugin.html
func (m *MiniProgram) GetLiveInfo(ctx context.Context, accessToken string, request GetLiveInfoRequest) (response *GetLiveInfoResponse, err error) {
	if err = m.Post(ctx, APIGetLiveInfo, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
