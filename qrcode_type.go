package wxm

type EnvVersion string

const (
	EnvVersionRelease EnvVersion = "release" // 正式版为
	EnvVersionTrial   EnvVersion = "trial"   // 体验版
	EnvVersionDevelop EnvVersion = "develop" // 开发版
)

// GetUnlimitedQRCodeParam https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
type GetUnlimitedQRCodeParam struct {
	Scene      string     `json:"scene"`
	Page       string     `json:"page"`
	Width      int        `json:"width,omitempty"`
	AutoColor  bool       `json:"auto_color"`
	LineColor  *LineColor `json:"line_color,omitempty"`
	IsHyaline  bool       `json:"is_hyaline"`
	EnvVersion EnvVersion `json:"env_version,omitempty"`
}

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type GetUnlimitedQRCodeResponse struct {
	Error
	Data []byte `json:"data"`
}
