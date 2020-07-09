package wxm

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	if this.signWithSHA1(ps) == signature {
		return true
	}
	return false
}

func (this *MiniProgram) signWithSHA1(ps []string) string {
	sort.Strings(ps)
	var s = sha1.New()
	s.Write([]byte(strings.Join(ps, "")))
	return hex.EncodeToString(s.Sum(nil))
}

// ParsePushMessage 小程序-获取来自微信服务器的推送消息 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Message_Encryption/Technical_Plan.html
func (this *MiniProgram) ParsePushMessage(token, timestamp, nonce, signature, key string, data []byte) (result *MessageInfo, err error) {
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	if result.Encrypt != "" {
		var ps = []string{token, timestamp, nonce, result.Encrypt}
		if this.signWithSHA1(ps) != signature {
			return nil, errors.New("failed to verify signature")
		}
		pData, err := this.decryptMessage(key, result.Encrypt)
		if err != nil {
			return nil, err
		}

		var index = bytes.LastIndex(pData, []byte("}"))
		pData = pData[20 : index+1]

		var info *MessageInfo
		if err = json.Unmarshal(pData, &info); err != nil {
			return nil, err
		}
		result = info
		return result, nil
	}
	return result, nil
}

func (this *MiniProgram) decryptMessage(key, encryptedData string) (result []byte, err error) {
	key = key + "="
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	decryptedBytes, err := AESCBCDecrypt(encryptedBytes, keyBytes, keyBytes[0:16])
	if err != nil {
		return nil, err
	}
	return decryptedBytes, err
}
