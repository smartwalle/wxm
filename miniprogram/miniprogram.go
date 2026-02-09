package miniprogram

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"sort"
	"strings"

	"github.com/smartwalle/wxm"
	"github.com/smartwalle/wxm/internal"
)

type MiniProgram struct {
	*wxm.Client
}

func New() *MiniProgram {
	return &MiniProgram{
		Client: wxm.New(),
	}
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

	plaintextBytes, err := internal.AESCBCDecrypt(ciphertextBytes, sessionKeyBytes, ivBytes)
	if err != nil {
		return nil, err
	}
	return plaintextBytes, err
}

// CheckMessageFromPushServer 验证消息来自微信服务器
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html#option-url
func (m *MiniProgram) CheckMessageFromPushServer(token, timestamp, nonce, signature string) bool {
	var values = []string{token, timestamp, nonce}
	return m.verifyMessage(values, signature)
}

// DecodePushMessage 获取来自微信服务器的推送消息
//
//	接口文档：https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Message_Encryption/Technical_Plan.html
func (m *MiniProgram) DecodePushMessage(token, timestamp, nonce, signature, key string, data []byte) (messageInfo *MessageInfo, err error) {
	if err = json.Unmarshal(data, &messageInfo); err != nil {
		return nil, err
	}

	if messageInfo.Encrypt != "" {
		var values = []string{token, timestamp, nonce, messageInfo.Encrypt}
		if !m.verifyMessage(values, signature) {
			return nil, errors.New("failed to verify signature")
		}
		var plaintext []byte
		plaintext, err = m.decryptMessage(key, messageInfo.Encrypt)
		if err != nil {
			return nil, err
		}

		var index = bytes.LastIndex(plaintext, []byte("}"))
		plaintext = plaintext[20 : index+1]

		var info *MessageInfo
		if err = json.Unmarshal(plaintext, &info); err != nil {
			return nil, err
		}
		messageInfo = info
		return messageInfo, nil
	}
	return messageInfo, nil
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

	plaintext, err := internal.AESCBCDecrypt(ciphertext, keyBytes, keyBytes[0:16])
	if err != nil {
		return nil, err
	}
	return plaintext, err
}
