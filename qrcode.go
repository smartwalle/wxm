package wxm

import (
	"encoding/json"
	"net/http"
)

const (
	kGetUnLimit = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

// GetUnlimited 小程序-获取小程序码 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (m *MiniProgram) GetUnlimited(param GetUnlimitedParam) (result *GetUnlimitedResponse, err error) {
	data, err := m.client.request(http.MethodPost, kGetUnLimit, param, nil, true, true)
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if err = json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
		return result, nil
	}

	result = &GetUnlimitedResponse{}
	result.Code = CodeSuccess
	result.Msg = "ok"
	result.Data = data

	return result, nil
}
