package wxm

import (
	"encoding/json"
	"net/http"
)

const (
	kGetLiveInfoURL = "https://api.weixin.qq.com/wxa/business/getliveinfo"
)

// GetLiveInfo 小程序-获取直播房间列表 https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/live-player-plugin.html
func (this *MiniProgram) GetLiveInfo(param GetLiveInfoParam) (result *GetLiveInfoRsp, err error) {
	data, err := this.client.requestWithAccessToken(http.MethodPost, kGetLiveInfoURL, param, nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
