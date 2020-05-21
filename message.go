package wxm

import (
	"encoding/json"
	"net/http"
)

const (
	kSendSubscribeMessageURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
	kSendTemplateMessageURL  = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

// SendSubscribeMessage 小程序-发送订阅消息 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (this *Client) SendSubscribeMessage(param SendSubscribeMessageParam) (result *SendSubscribeMessageRsp, err error) {
	data, err := this.request(http.MethodPost, kSendSubscribeMessageURL, param, nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SendTemplateMessage 公众号-发送模板消息 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (this *Client) SendTemplateMessage(param SendTemplateMessageParam) (result *SendTemplateMessageRsp, err error) {
	data, err := this.request(http.MethodPost, kSendTemplateMessageURL, param, nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
