package wxm

import (
	"context"
	"net/http"
)

const (
	kSendSubscribeMessage       = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
	kSendUniformMessage         = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
	kSendCustomerServiceMessage = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	kSendTemplateMessage        = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)

// SendSubscribeMessage 小程序-发送订阅消息 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (m *MiniProgram) SendSubscribeMessage(ctx context.Context, param SendSubscribeMessageParam) (result *SendSubscribeMessageResponse, err error) {
	if err = m.client.requestWithAccessToken(ctx, http.MethodPost, kSendSubscribeMessage, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SendUniformMessage 小程序-下发小程序和公众号统一的服务消息 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (m *MiniProgram) SendUniformMessage(ctx context.Context, param SendUniformMessageParam) (result *SendUniformMessageReponse, err error) {
	if err = m.client.requestWithAccessToken(ctx, http.MethodPost, kSendUniformMessage, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SendCustomerServiceMessage 小程序-发送客服消息给用户 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func (m *MiniProgram) SendCustomerServiceMessage(ctx context.Context, param SendCustomerServiceMessageParam) (result *SendCustomerServiceMessageResponse, err error) {
	if err = m.client.requestWithAccessToken(ctx, http.MethodPost, kSendCustomerServiceMessage, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SendTemplateMessage 公众号-发送模板消息 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (o *OfficialAccount) SendTemplateMessage(ctx context.Context, param SendTemplateMessageParam) (result *SendTemplateMessageResponse, err error) {
	if err = o.client.requestWithAccessToken(ctx, http.MethodPost, kSendTemplateMessage, param, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
