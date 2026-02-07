package miniprogram

import (
	"context"
	"encoding/json"
)

const (
	kGetUserPhoneNumber = "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
)

// GetUserPhoneNumber 获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func (m *MiniProgram) GetUserPhoneNumber(ctx context.Context, accessToken, code string) (response *GetUserPhoneNumberResponse, err error) {
	var request = struct {
		Code string `json:"code"`
	}{
		Code: code,
	}
	if err = m.Post(ctx, accessToken, kGetUserPhoneNumber, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// DecodePhoneNumber 解密手机号码数据 https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/deprecatedGetPhoneNumber.html
//
//	小程序端申请获取用户的手机号码之后，获取到的是加密的数据，需要调用本方法对该数据进行解密，以获取手机号码。
func (m *MiniProgram) DecodePhoneNumber(sessionKey, encryptedData, iv string) (phoneInfo *PhoneInfo, err error) {
	plaintext, err := m.decrypt(sessionKey, encryptedData, iv)
	if err = json.Unmarshal(plaintext, &phoneInfo); err != nil {
		return nil, err
	}
	return phoneInfo, nil
}

// DecodeUserInfo 解密用户数据 https://developers.weixin.qq.com/miniprogram/dev/api/open-api/user-info/wx.getUserInfo.html
//
//	小程序端申请获取用户的信息之后，获取到的有加密的数据，需要调用本方法对该数据进行解密，以获取加密信息。
func (m *MiniProgram) DecodeUserInfo(sessionKey, encryptedData, iv string) (response *UserInfoResponse, err error) {
	plaintext, err := m.decrypt(sessionKey, encryptedData, iv)
	if err = json.Unmarshal(plaintext, &response); err != nil {
		return nil, err
	}
	return response, nil
}
