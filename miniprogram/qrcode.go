package miniprogram

import (
	"context"
	"encoding/json"

	"github.com/smartwalle/wxm"
)

const (
	kGetUnLimit = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

// GetUnlimitedQRCode 小程序-获取小程序码
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (m *MiniProgram) GetUnlimitedQRCode(ctx context.Context, accessToken string, request GetUnlimitedQRCodeRequest) (response *GetUnlimitedQRCodeResponse, err error) {
	data, err := m.Request(ctx, accessToken, kGetUnLimit, request, nil)
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if err = json.Unmarshal(data, &response); err != nil {
			return nil, err
		}
		return response, nil
	}

	response = &GetUnlimitedQRCodeResponse{}
	response.Code = wxm.CodeSuccess
	response.Msg = "ok"
	response.Data = data

	return response, nil
}
