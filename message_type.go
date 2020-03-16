package wxm

type MessageData map[string]map[string]string

func (this MessageData) add(key, value string) {
	var m = make(map[string]string)
	m["value"] = value
	this[key] = m
}

type SendMessageParam struct {
	ToUser           string      `json:"touser"`                      // 是 接收者（用户）的 openid
	TemplateId       string      `json:"template_id"`                 // 是 所需下发的订阅模板id
	Data             MessageData `json:"data"`                        // 是 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
	Page             string      `json:"page,omitempty"`              // 否 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	MiniProgramState string      `json:"miniprogram_state,omitempty"` // 否 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string      `json:"lang,omitempty"`              // 否 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

func (this *SendMessageParam) AddData(key, value string) {
	if this.Data == nil {
		this.Data = make(MessageData)
	}
	this.Data.add(key, value)
}

type SendMessageRsp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
}
