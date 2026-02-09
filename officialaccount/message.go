package officialaccount

import (
	"context"
)

const (
	kSendTemplateMessage = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)

// SendTemplateMessage 发送模板消息
//
//	接口文档：https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (o *OfficialAccount) SendTemplateMessage(ctx context.Context, accessToken string, request SendTemplateMessageRequest) (response *SendTemplateMessageResponse, err error) {
	if err = o.Post(ctx, accessToken, kSendTemplateMessage, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
