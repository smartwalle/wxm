package wxm

import (
	"net/http"
)

const (
	kGetLiveInfo = "https://api.weixin.qq.com/wxa/business/getliveinfo"
)

// GetLiveInfo 小程序-获取直播房间列表 https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/live-player-plugin.html
func (this *MiniProgram) GetLiveInfo(param GetLiveInfoParam) (result *GetLiveInfoRsp, err error) {
	if err = this.client.requestWithAccessToken(http.MethodPost, kGetLiveInfo, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
