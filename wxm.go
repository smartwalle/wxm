package wxm

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/smartwalle/ngx"
)

const (
	kDomain     = "https://api.weixin.qq.com"
	APIGetToken = "/cgi-bin/token"
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
	var v = url.Values{}
	v.Add("appid", appId)
	v.Add("secret", secret)
	v.Add("grant_type", "client_credential")
	if err = c.Get(ctx, APIGetToken, "", v, &token); err != nil {
		return nil, err
	}
	return token, nil
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
