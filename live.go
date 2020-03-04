package wxm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	kGetLiveInfoURL = "http://api.weixin.qq.com/wxa/business/getliveinfo?access_token=%s"
)

// GetLiveInfo 获取直播房间列表
func (this *Client) GetLiveInfo(param GetLiveInfoParam) (result *GetLiveInfoRsp, err error) {
	accessToken, err := this.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf(kGetLiveInfoURL, accessToken)

	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	rsp, err := this.client.Do(req)
	if rsp != nil && rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
