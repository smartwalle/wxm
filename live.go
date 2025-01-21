package wxm

import (
	"context"
	"net/http"
)

const (
	kGetLiveInfo = "https://api.weixin.qq.com/wxa/business/getliveinfo"
)

// GetLiveInfo 小程序-获取直播房间列表 https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/live-player-plugin.html
func (m *MiniProgram) GetLiveInfo(ctx context.Context, param GetLiveInfoParam) (result *GetLiveInfoResponse, err error) {
	if err = m.client.requestWithAccessToken(ctx, http.MethodPost, kGetLiveInfo, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
