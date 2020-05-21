package wxm

import (
	"encoding/json"
	"net/http"
)

const (
	kGetLiveInfoURL = "http://api.weixin.qq.com/wxa/business/getliveinfo?access_token=%s"
)

// GetLiveInfo 小程序-获取直播房间列表 https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/live-player-plugin.html
func (this *Client) GetLiveInfo(param GetLiveInfoParam) (result *GetLiveInfoRsp, err error) {
	data, err := this.request(http.MethodPost, kGetLiveInfoURL, param, nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
