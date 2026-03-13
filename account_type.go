package wxm

// GetAccountBasicInfoResponse 获取基本信息响应
type GetAccountBasicInfoResponse struct {
	Error
	AppId             string         `json:"appid"`              // 账号 appid
	AccountType       int            `json:"account_type"`       // 账号类型（1：订阅号，2：服务号，3：小程序）
	PrincipalType     int            `json:"principal_type"`     // 主体类型
	PrincipalName     string         `json:"principal_name"`     // 主体名称
	RealNameStatus    int            `json:"realname_status"`    // 实名验证状态
	WXVerifyInfo      *WXVerifyInfo  `json:"wx_verify_info"`     // 微信认证信息
	SignatureInfo     *SignatureInfo `json:"signature_info"`     // 功能介绍信息
	HeadImageInfo     *HeadImageInfo `json:"head_image_info"`    // 头像信息
	Nickname          string         `json:"nickname"`           // 小程序名称
	RegisteredCountry int            `json:"registered_country"` // 注册国家
	NicknameInfo      *NicknameInfo  `json:"nickname_info"`      // 名称信息
	Credential        string         `json:"credential"`         // 非个人主体时返回的是企业或者政府或其他组织的代号
	CustomerType      int            `json:"customer_type"`      // 认证类型；如果未完成微信认证则返回0
}

// WXVerifyInfo 微信认证信息
type WXVerifyInfo struct {
	QualificationVerify   bool  `json:"qualification_verify"`     // 是否资质认证，若是，拥有微信认证相关的权限
	NamingVerify          bool  `json:"naming_verify"`            // 是否名称认证
	AnnualReview          bool  `json:"annual_review"`            // 是否需要年审（qualification_verify == true 时才有该字段）
	AnnualReviewBeginTime int64 `json:"annual_review_begin_time"` // 年审开始时间，时间戳（qualification_verify == true 时才有该字段）
	AnnualReviewEndTime   int64 `json:"annual_review_end_time"`   // 年审截止时间，时间戳（qualification_verify == true 时才有该字段）
}

// SignatureInfo 功能介绍信息
type SignatureInfo struct {
	Signature       string `json:"signature"`         // 功能介绍
	ModifyUsedCount int    `json:"modify_used_count"` // 功能介绍已使用修改次数（本月）
	ModifyQuota     int    `json:"modify_quota"`      // 功能介绍修改次数总额度（本月）
}

// HeadImageInfo 头像信息
type HeadImageInfo struct {
	HeadImageURL    string `json:"head_image_url"`    // 头像 url
	ModifyUsedCount int    `json:"modify_used_count"` // 头像已使用修改次数（本年）
	ModifyQuota     int    `json:"modify_quota"`      // 头像修改次数总额度（本年）
}

// NicknameInfo 名称信息
type NicknameInfo struct {
	Nickname        string `json:"nickname"`          // 小程序/公众号账号名称
	ModifyUsedCount int    `json:"modify_used_count"` // 小程序名称已使用修改次数（本年）
	ModifyQuota     int    `json:"modify_quota"`      // 小程序名称修改次数总额度（本年）
}
