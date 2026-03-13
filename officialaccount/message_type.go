package officialaccount

import (
	"github.com/smartwalle/wxm"
)

type MessageData map[string]map[string]string

func (m MessageData) add(param, key, value string) {
	var values = m[param]
	if values == nil {
		values = make(map[string]string)
	}
	values[key] = value
	m[param] = values
}

type MiniProgramInfo struct {
	AppId    string `json:"appid"`    // 是 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagepath"` // 否 所需跳转到小程序的具体页面路径，支持带参数, （示例index?foo = bar），要求该小程序已发布，暂不支持小游戏
}

func NewMiniProgramInfo(appId, pagePath string) *MiniProgramInfo {
	return &MiniProgramInfo{
		AppId:    appId,
		PagePath: pagePath,
	}
}

// SendTemplateMessageRequest https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
type SendTemplateMessageRequest struct {
	ToUser      string           `json:"touser"`                // 是 接收者（用户）的 openid
	TemplateId  string           `json:"template_id"`           // 是 模板Id
	URL         string           `json:"url"`                   // 否 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MiniProgramInfo `json:"miniprogram,omitempty"` // 否 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        MessageData      `json:"data"`                  // 是 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
}

func (m *SendTemplateMessageRequest) AddData(key, value, color string) {
	if m.Data == nil {
		m.Data = make(MessageData)
	}
	m.Data.add(key, "value", value)
	m.Data.add(key, "color", color)
}

type SendTemplateMessageResponse struct {
	wxm.Error
}
