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
	*client
}

func NewMiniProgram(appId, appSecret string) *MiniProgram {
	var c = &MiniProgram{}
	c.client = newClient(appId, appSecret)
	return c
}

// GetToken 小程序-获取全局唯一后台接口调用凭据（access_token）https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (m *MiniProgram) GetToken() (result string, err error) {
	return m.client.GetToken()
}

// RefreshToken 小程序-刷新本地全局唯一后台接口调用凭据（access_token）
func (m *MiniProgram) RefreshToken() (err error) {
	return m.client.RefreshToken()
}

func (m *MiniProgram) decrypt(sessionKey, ciphertext, iv string) (result []byte, err error) {
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	plaintextBytes, err := AESCBCDecrypt(ciphertextBytes, sessionKeyBytes, ivBytes)
	if err != nil {
		return nil, err
	}
	return plaintextBytes, err
}

// CheckMessageFromPushServer 小程序-验证消息来自微信服务器 https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html#option-url
func (m *MiniProgram) CheckMessageFromPushServer(token, timestamp, nonce, signature string) bool {
	var nList = []string{token, timestamp, nonce}
	return m.verifyMessage(nList, signature)
}

// DecodePushMessage 小程序-获取来自微信服务器的推送消息 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Message_Encryption/Technical_Plan.html
func (m *MiniProgram) DecodePushMessage(token, timestamp, nonce, signature, key string, data []byte) (result *MessageInfo, err error) {
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	if result.Encrypt != "" {
		var nList = []string{token, timestamp, nonce, result.Encrypt}
		if !m.verifyMessage(nList, signature) {
			return nil, errors.New("failed to verify signature")
		}
		var plaintext []byte
		plaintext, err = m.decryptMessage(key, result.Encrypt)
		if err != nil {
			return nil, err
		}

		var index = bytes.LastIndex(plaintext, []byte("}"))
		plaintext = plaintext[20 : index+1]

		var info *MessageInfo
		if err = json.Unmarshal(plaintext, &info); err != nil {
			return nil, err
		}
		result = info
		return result, nil
	}
	return result, nil
}

func (m *MiniProgram) verifyMessage(values []string, signature string) bool {
	sort.Strings(values)
	var hashed = sha1.New()
	hashed.Write([]byte(strings.Join(values, "")))
	return hex.EncodeToString(hashed.Sum(nil)) == signature
}

func (m *MiniProgram) decryptMessage(key, data string) (result []byte, err error) {
	key = key + "="
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	plaintext, err := AESCBCDecrypt(ciphertext, keyBytes, keyBytes[0:16])
	if err != nil {
		return nil, err
	}
	return plaintext, err
}
