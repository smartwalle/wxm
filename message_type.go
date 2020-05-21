package wxm

type MiniProgramState string

const (
	MiniProgramStateDeveloper MiniProgramState = "developer" // 为开发版
	MiniProgramStateTrial     MiniProgramState = "trial"     // 为体验版
	MiniProgramStateFormal    MiniProgramState = "formal"    // 为正式版
)

type MessageData map[string]map[string]string

func (this MessageData) add(param, key, value string) {
	var m = this[param]
	if m == nil {
		m = make(map[string]string)
	}
	m[key] = value
	this[param] = m
}

// SendSubscribeMessageParam https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
type SendSubscribeMessageParam struct {
	ToUser           string           `json:"touser"`                      // 是 接收者（用户）的 openid
	TemplateId       string           `json:"template_id"`                 // 是 所需下发的订阅模板id
	Data             MessageData      `json:"data"`                        // 是 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
	Page             string           `json:"page,omitempty"`              // 否 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	MiniProgramState MiniProgramState `json:"miniprogram_state,omitempty"` // 否 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string           `json:"lang,omitempty"`              // 否 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

func (this *SendSubscribeMessageParam) AddData(key, value string) {
	if this.Data == nil {
		this.Data = make(MessageData)
	}
	this.Data.add(key, "value", value)
}

type SendSubscribeMessageRsp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
}

type MiniProgram struct {
	AppId    string `json:"appid"`    // 是 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagepath"` // 否 所需跳转到小程序的具体页面路径，支持带参数, （示例index?foo = bar），要求该小程序已发布，暂不支持小游戏
}

// SendTemplateMessageParam https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
type SendTemplateMessageParam struct {
	ToUser      string       `json:"touser"`                // 是 接收者（用户）的 openid
	TemplateId  string       `json:"template_id"`           // 是 模板ID
	URL         string       `json:"url"`                   // 否 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"` // 否 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        MessageData  `json:"data"`                  // 是 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
}

func (this *SendTemplateMessageParam) AddData(key, value, color string) {
	if this.Data == nil {
		this.Data = make(MessageData)
	}
	this.Data.add(key, "value", value)
	this.Data.add(key, "color", color)
}

type SendTemplateMessageRsp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
}
