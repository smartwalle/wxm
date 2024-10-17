package wxm

type MiniProgramState string

const (
	MiniProgramStateDeveloper MiniProgramState = "developer" // 为开发版
	MiniProgramStateTrial     MiniProgramState = "trial"     // 为体验版
	MiniProgramStateFormal    MiniProgramState = "formal"    // 为正式版
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

// SendSubscribeMessage https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
type SendSubscribeMessage struct {
	ToUser           string           `json:"touser"`                      // 是 接收者（用户）的 openid
	TemplateId       string           `json:"template_id"`                 // 是 所需下发的订阅模板id
	Data             MessageData      `json:"data"`                        // 是 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
	Page             string           `json:"page,omitempty"`              // 否 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	MiniProgramState MiniProgramState `json:"miniprogram_state,omitempty"` // 否 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string           `json:"lang,omitempty"`              // 否 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

func (m *SendSubscribeMessage) AddData(key, value string) {
	if m.Data == nil {
		m.Data = make(MessageData)
	}
	m.Data.add(key, "value", value)
}

type SendSubscribeMessageRsp struct {
	Error
}

// SendUniformMessage https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
type SendUniformMessage struct {
	ToUser           string            `json:"touser"`                       // 是 用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
	WeAppTemplateMsg *WeAppTemplateMsg `json:"weapp_template_msg,omitempty"` // 否 小程序模板消息相关的信息，可以参考小程序模板消息接口; 有此节点则优先发送小程序模板消息
	MPTemplateMsg    *MPTemplateMsg    `json:"mp_template_msg,omitempty"`    // 是 公众号模板消息相关的信息，可以参考公众号模板消息接口；有此节点并且没有weapp_template_msg节点时，发送公众号模板消息
}

type WeAppTemplateMsg struct {
	TemplateId      string      `json:"template_id"`      // 是	小程序模板ID
	Page            string      `json:"page"`             // 是	小程序页面路径
	FormId          string      `json:"form_id"`          // 是	小程序模板消息 formid
	Data            MessageData `json:"data"`             // 是	小程序模板数据
	EmphasisKeyword string      `json:"emphasis_keyword"` // 是	小程序模板放大关键词
}

func (m *WeAppTemplateMsg) AddData(key, value string) {
	if m.Data == nil {
		m.Data = make(MessageData)
	}
	m.Data.add(key, "value", value)
}

type MPTemplateMsg struct {
	AppId       string           `json:"appid"`                 // 是 公众号appid，要求与小程序有绑定且同主体
	TemplateId  string           `json:"template_id"`           // 是 模板ID
	URL         string           `json:"url"`                   // 否 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MiniProgramInfo `json:"miniprogram,omitempty"` // 否 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        MessageData      `json:"data"`                  // 是 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
}

func (m *MPTemplateMsg) AddData(key, value, color string) {
	if m.Data == nil {
		m.Data = make(MessageData)
	}
	m.Data.add(key, "value", value)
	m.Data.add(key, "color", color)
}

type SendUniformMessageRsp struct {
	Error
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

// SendCustomerServiceMessage https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
type SendCustomerServiceMessage struct {
	ToUser          string              `json:"touser"`                    // 是 用户的 OpenID
	MsgType         MsgType             `json:"msgtype"`                   // 是 消息类型
	Text            *MsgText            `json:"text,omitempty"`            // 是 文本消息，msgtype="text" 时必填
	Image           *MsgImage           `json:"image,omitempty"`           // 是 图片消息，msgtype="image" 时必填
	Link            *MsgLink            `json:"link,omitempty"`            // 是 图文链接，msgtype="link" 时必填
	MiniProgramPage *MsgMiniProgramPage `json:"miniprogrampage,omitempty"` // 是 小程序卡片，msgtype="miniprogrampage" 时必填

}

type MsgType string

const (
	MsgTypeOfText            MsgType = "text"
	MsgTypeOfImage           MsgType = "image"
	MsgTypeOfLink            MsgType = "link"
	MsgTypeOfMiniProgramPage MsgType = "miniprogrampage"
)

type MsgText struct {
	Content string `json:"content"`
}

type MsgImage struct {
	MediaId string `json:"media_id"`
}

type MsgLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}

type MsgMiniProgramPage struct {
	Title        string `json:"title"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

type SendCustomerServiceMessageRsp struct {
	Error
}

// SendTemplateMessage https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
type SendTemplateMessage struct {
	ToUser      string           `json:"touser"`                // 是 接收者（用户）的 openid
	TemplateId  string           `json:"template_id"`           // 是 模板ID
	URL         string           `json:"url"`                   // 否 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MiniProgramInfo `json:"miniprogram,omitempty"` // 否 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        MessageData      `json:"data"`                  // 是 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
}

func (m *SendTemplateMessage) AddData(key, value, color string) {
	if m.Data == nil {
		m.Data = make(MessageData)
	}
	m.Data.add(key, "value", value)
	m.Data.add(key, "color", color)
}

type SendTemplateMessageRsp struct {
	Error
}
