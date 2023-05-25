package wxm

// GetUnlimited https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
type GetUnlimited struct {
	Scene     string     `json:"scene"`
	Page      string     `json:"page"`
	Width     int        `json:"width,omitempty"`
	AutoColor bool       `json:"auto_color"`
	LineColor *LineColor `json:"line_color,omitempty"`
	IsHyaline bool       `json:"is_hyaline"`
}

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type GetUnlimitedRsp struct {
	Error
	Data []byte `json:"data"`
}
