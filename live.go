package wxm

import (
	"encoding/json"
	"net/http"
)

const (
	kGetLiveInfoURL = "http://api.weixin.qq.com/wxa/business/getliveinfo?access_token=%s"
)

// GetLiveInfo 获取直播房间列表
func (this *Client) GetLiveInfo(param GetLiveInfoParam) (result *GetLiveInfoRsp, err error) {
	data, err := this.request(http.MethodPost, kGetLiveInfoURL, param)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
