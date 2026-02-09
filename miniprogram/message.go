package miniprogram

import (
	"context"
)

const (
	APISendSubscribeMessage       = "/cgi-bin/message/subscribe/send"
	APISendUniformMessage         = "/cgi-bin/message/wxopen/template/uniform_send"
	APISendCustomerServiceMessage = "/cgi-bin/message/custom/send"
)

// SendSubscribeMessage 发送订阅消息
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (m *MiniProgram) SendSubscribeMessage(ctx context.Context, accessToken string, request SendSubscribeMessageRequest) (response *SendSubscribeMessageResponse, err error) {
	if err = m.Post(ctx, APISendSubscribeMessage, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// SendUniformMessage 下发小程序和公众号统一的服务消息
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (m *MiniProgram) SendUniformMessage(ctx context.Context, accessToken string, request SendUniformMessageRequest) (response *SendUniformMessageReponse, err error) {
	if err = m.Post(ctx, APISendUniformMessage, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// SendCustomerServiceMessage 发送客服消息给用户
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func (m *MiniProgram) SendCustomerServiceMessage(ctx context.Context, accessToken string, request SendCustomerServiceMessageRequest) (response *SendCustomerServiceMessageResponse, err error) {
	if err = m.Post(ctx, APISendCustomerServiceMessage, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
