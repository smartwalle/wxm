package wxm

import "encoding/base64"

type MiniProgram struct {
	client *client
}

func NewMiniProgram(appId, appSecret string) *MiniProgram {
	var c = &MiniProgram{}
	c.client = newClient(appId, appSecret)
	return c
}

// GetToken 小程序-获取全局唯一后台接口调用凭据（access_token）https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (this *MiniProgram) GetToken() (result string, err error) {
	return this.client.GetToken()
}

// RefreshToken 小程序-刷新本地全局唯一后台接口调用凭据（access_token）
func (this *MiniProgram) RefreshToken() (err error) {
	return this.client.RefreshToken()
}

func (this *MiniProgram) decrypt(sessionKey, encryptedData, iv string) (result []byte, err error) {
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	decryptedBytes, err := AESCBCDecrypt(encryptedBytes, sessionKeyBytes, ivBytes)
	if err != nil {
		return nil, err
	}
	return decryptedBytes, err
}
