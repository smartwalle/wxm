package wxm

type MessageInfo struct {
	ToUserName   string  `json:"ToUserName"`
	Encrypt      string  `json:"Encrypt"`
	FromUserName string  `json:"FromUserName"`
	CreateTime   int64   `json:"CreateTime"`
	MsgType      MsgType `json:"MsgType"`
	MsgId        int64   `json:"MsgId"`
	Content      string  `json:"content"`
	PicURL       string  `json:"PicUrl"`
	MediaId      string  `json:"MediaId"`
	Title        string  `json:"Title"`
	AppId        string  `json:"AppId"`
	PagePath     string  `json:"PagePath"`
	ThumbUrl     string  `json:"ThumbUrl"`
	ThumbMediaId string  `json:"ThumbMediaId"`
}
