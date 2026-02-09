package wxm

import (
	"fmt"
)

type Code int

func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}

func (c Code) IsFailure() bool {
	return c != CodeSuccess
}

const (
	MessageSuccess = "ok"
)

const (
	CodeSuccess           Code = 0     // 请求成功
	CodeInvalidCredential Code = 40001 // access_token 无效或者 AppSecret 错误
	CodeInvalidGrantType  Code = 40002 // 请确保 grant_type 字段值为 client_credential
	CodeInvalidAppId      Code = 40013 // 不合法的 AppID，请开发者检查 AppID 的正确性，避免异常字符，注意大小写
	CodeInvalidCode       Code = 40029 // 不合法的 code
	CodeMaxRate           Code = 45009 // 调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成。
	CodeInvalidPage       Code = 41030 // 所传page页面不存在，或者小程序没有发布
)

type Error struct {
	Code Code   `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (err Error) Error() string {
	return fmt.Sprintf("%d-%s", err.Code, err.Msg)
}

func (err Error) IsSuccess() bool {
	return err.Code.IsSuccess()
}

func (err Error) IsFailure() bool {
	return err.Code.IsFailure()
}

type Token struct {
	Error
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// RIDRequest RID请求详情
type RIDRequest struct {
	InvokeTime   int64  `json:"invoke_time"`   // 发起请求的时间戳
	CostInMs     int64  `json:"cost_in_ms"`    // 请求毫秒级耗时
	RequestURL   string `json:"request_url"`   // 请求的URL参数
	RequestBody  string `json:"request_body"`  // post请求的请求参数
	ResponseBody string `json:"response_body"` // 接口请求返回参数
	ClientIP     string `json:"client_ip"`     // 接口请求的客户端ip
}

// GetRIDInfoResponse 查询RID信息响应
type GetRIDInfoResponse struct {
	Error
	Request *RIDRequest `json:"request"` // 该rid对应的请求详情
}

// Quota API调用额度详情，表示某个接口的调用配额信息
type Quota struct {
	DailyLimit int64 `json:"daily_limit"` // 当天该账号可调用该接口的次数
	Used       int64 `json:"used"`        // 当天已经调用的次数
	Remain     int64 `json:"remain"`      // 当天剩余调用次数
}

// RateLimit 调用频率限制，表示接口在某个周期内的调用限制
type RateLimit struct {
	CallCount     int64 `json:"call_count"`     // 周期内可调用数量，单位 次
	RefreshSecond int64 `json:"refresh_second"` // 更新周期，单位 秒
}

// GetAPIQuotaResponse 查询API调用额度响应
type GetAPIQuotaResponse struct {
	Error
	Quota              *Quota     `json:"quota"`                // quota详情
	RateLimit          *RateLimit `json:"rate_limit"`           // 普通调用频率限制
	ComponentRateLimit *RateLimit `json:"component_rate_limit"` // 代调用频率限制
}

// ClearQuotaResponse 重置API调用次数响应
type ClearQuotaResponse struct {
	Error
}

// ClearAPIQuotaResponse 重置指定API调用次数响应
type ClearAPIQuotaResponse struct {
	Error
}
