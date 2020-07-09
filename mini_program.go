package wxm

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"sort"
	"strings"
)

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

// CheckMessagePushServer 小程序-验证消息来自微信服务器 https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html#option-url
func (this *MiniProgram) CheckMessagePushServer(token, timestamp, nonce, signature string) bool {
	var ps = []string{token, timestamp, nonce}
	sort.Strings(ps)

	var s = sha1.New()
	s.Write([]byte(strings.Join(ps, "")))
	if hex.EncodeToString(s.Sum(nil)) == signature {
		return true
	}
	return false
}
