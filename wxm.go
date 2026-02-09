package wxm

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/smartwalle/ngx"
)

const (
	kDomain           = "https://api.weixin.qq.com"
	APIGetToken       = "/cgi-bin/token"
	APIGetStableToken = "/cgi-bin/stable_token"
	APIGetRIDInfo     = "/cgi-bin/openapi/rid/get"
	APIGetAPIQuota    = "/cgi-bin/openapi/quota/get"
	APIClearQuota     = "/cgi-bin/clear_quota"
	APIClearQuotaV2   = "/cgi-bin/clear_quota/v2"
	APIClearAPIQuota  = "/cgi-bin/openapi/quota/clear"
)

type Client struct {
	HTTPClient *http.Client
}

func New() *Client {
	var c = &Client{}
	c.HTTPClient = http.DefaultClient
	return c
}

// GetToken 获取全局唯一后台接口调用凭据
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *Client) GetToken(ctx context.Context, appId, secret string) (token *Token, err error) {
	var query = url.Values{}
	query.Add("appid", appId)
	query.Add("secret", secret)
	query.Add("grant_type", "client_credential")
	if err = c.Get(ctx, APIGetToken, "", query, &token); err != nil {
		return nil, err
	}
	return token, nil
}

// GetStableToken 获取稳定版接口调用凭据
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/mp-access-token/api_getstableaccesstoken.html
func (c *Client) GetStableToken(ctx context.Context, appId, secret string, forceRefresh bool) (token *Token, err error) {
	var request = map[string]interface{}{
		"appid":         appId,
		"secret":        secret,
		"grant_type":    "client_credential",
		"force_refresh": forceRefresh,
	}
	if err = c.Post(ctx, APIGetStableToken, "", request, nil, &token); err != nil {
		return nil, err
	}
	return token, nil
}

// GetRIDInfo 查询RID信息
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/openApi-mgnt/api_getridinfo.html
func (c *Client) GetRIDInfo(ctx context.Context, accessToken, rid string) (response *GetRIDInfoResponse, err error) {
	var request = map[string]string{
		"rid": rid,
	}
	if err = c.Post(ctx, APIGetRIDInfo, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetAPIQuota 查询API调用额度
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/openApi-mgnt/api_getapiquota.html
func (c *Client) GetAPIQuota(ctx context.Context, accessToken, cgiPath string) (response *GetAPIQuotaResponse, err error) {
	var request = map[string]string{
		"cgi_path": cgiPath,
	}
	if err = c.Post(ctx, APIGetAPIQuota, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// ClearQuota 重置API调用次数
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/openApi-mgnt/api_clearquota.html
func (c *Client) ClearQuota(ctx context.Context, accessToken, appId string) (response *ClearQuotaResponse, err error) {
	var request = map[string]string{
		"appid": appId,
	}
	if err = c.Post(ctx, APIClearQuota, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// ClearQuotaV2 使用AppSecret重置API调用次数
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/openApi-mgnt/api_clearquotabyappsecret.html
func (c *Client) ClearQuotaV2(ctx context.Context, appId, appSecret string) (response *ClearQuotaResponse, err error) {
	var request = map[string]string{
		"appid":     appId,
		"appsecret": appSecret,
	}

	if err = c.Post(ctx, APIClearQuotaV2, "", request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// ClearAPIQuota 重置指定API调用次数
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/openApi-mgnt/api_clearapiquota.html
func (c *Client) ClearAPIQuota(ctx context.Context, accessToken, cgiPath string) (response *ClearAPIQuotaResponse, err error) {
	var request = map[string]string{
		"cgi_path": cgiPath,
	}
	if err = c.Post(ctx, APIClearAPIQuota, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) buildRequest(method, api, accessToken string, query url.Values) *ngx.Request {
	if query == nil {
		query = url.Values{}
	}
	if accessToken != "" {
		query.Set("access_token", accessToken)
	}

	var req = ngx.NewRequest(method, kDomain)
	req.JoinPath(api)
	req.Client = c.HTTPClient
	req.Query = query
	return req
}

func (c *Client) Post(ctx context.Context, api, accessToken string, payload interface{}, query url.Values, dst interface{}) (err error) {
	var req = c.buildRequest(ngx.Post, api, accessToken, query)
	if payload != nil {
		req.Body = ngx.JSONEncoder(payload)
	}

	if _, err = req.Decode(ctx, ngx.JSONDecoder(&dst)); err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(ctx context.Context, api, accessToken string, query url.Values, dst interface{}) (err error) {
	var req = c.buildRequest(ngx.Get, api, accessToken, query)
	if _, err = req.Decode(ctx, ngx.JSONDecoder(&dst)); err != nil {
		return err
	}
	return nil
}

func (c *Client) Request(ctx context.Context, method, api, accessToken string, payload interface{}, query url.Values) (result []byte, err error) {
	var req = c.buildRequest(method, api, accessToken, query)
	if payload != nil {
		req.Body = ngx.JSONEncoder(payload)
	}

	resp, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Upload(ctx context.Context, api, accessToken, fieldname, filename, filepath string, query url.Values, dst interface{}) (err error) {
	var req = c.buildRequest(ngx.Post, api, accessToken, query)
	req.FileForm.AddFilePath(fieldname, filename, filepath)

	if _, err = req.Decode(ctx, ngx.JSONDecoder(&dst)); err != nil {
		return err
	}
	return nil
}
