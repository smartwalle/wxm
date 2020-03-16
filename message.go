package wxm

import (
	"encoding/json"
	"net/http"
)

const (
	kSendMessageURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
)

// SendMessage 发送消息
func (this *Client) SendMessage(param SendMessageParam) (result *SendMessageRsp, err error) {
	data, err := this.request(http.MethodPost, kSendMessageURL, param)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
