package wxm

import (
	"encoding/json"
	"net/http"
)

const (
	kGetUnLimitURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)

// GetUnlimited 小程序-获取小程序码 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (this *MiniProgram) GetUnlimited(param GetUnlimitedParam) (result *GetUnlimitedRsp, err error) {
	data, err := this.client.request(http.MethodPost, kGetUnLimitURL, param, nil)
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if err = json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
		return result, nil
	}

	result = &GetUnlimitedRsp{}
	result.ErrCode = 0
	result.ErrMsg = "ok"
	result.Data = data

	return result, nil
}
