package wxm

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	kGetToken = "https://api.weixin.qq.com/cgi-bin/token"
)

type Option func(client *client)

func WithAccessToken(accessToken string) Option {
	return func(c *client) {
		c.accessToken = accessToken
	}
}

func WithHTTPClient(c *http.Client) Option {
	return func(nc *client) {
		nc.client = c
	}
}

type client struct {
	appId       string
	appSecret   string
	client      *http.Client
	accessToken string
}

func newClient(appId, appSecret string) *client {
	var c = &client{}
	c.appId = appId
	c.appSecret = appSecret
	c.client = http.DefaultClient
	return c
}

func (c *client) With(opts ...Option) *client {
	var n = *c
	var np = &n
	for _, opt := range opts {
		if opt != nil {
			opt(np)
		}
	}
	return np
}

func (c *client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

func (c *client) SetHTTPClient(client *http.Client) {
	c.client = client
}

// GetToken 小程序、公众号-获取全局唯一后台接口调用凭据（access_token） https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *client) GetToken() (token *Token, err error) {
	var values = url.Values{}
	values.Add("appid", c.appId)
	values.Add("secret", c.appSecret)
	values.Add("grant_type", "client_credential")

	if err = c.requestWithoutAccessToken(http.MethodGet, kGetToken, nil, values, &token); err != nil {
		return nil, err
	}

	if token != nil && token.IsSuccess() {
		token.CreateTime = time.Now().Unix()
	}

	return token, nil
}

func (c *client) requestWithAccessToken(method, api string, param interface{}, values url.Values, result interface{}) error {
	var data, err = c.request(method, api, true, param, values)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}

func (c *client) requestWithoutAccessToken(method, api string, param interface{}, values url.Values, result interface{}) error {
	var data, err = c.request(method, api, false, param, values)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}

func (c *client) request(method, api string, needAuth bool, param interface{}, values url.Values) (result []byte, err error) {
	if values == nil {
		values = url.Values{}
	}

	if needAuth {
		values.Set("access_token", c.accessToken)
	}

	var body io.Reader
	if param != nil {
		data, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(data)
	}

	var nURL = api + "?" + values.Encode()
	req, err := http.NewRequest(method, nURL, body)
	if err != nil {
		return nil, err
	}
	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	result, err = io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) upload(method, api, fieldname, filename string, values url.Values, result interface{}) error {
	if values == nil {
		values = url.Values{}
	}
	values.Set("access_token", c.accessToken)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldname, filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	if err = writer.Close(); err != nil {
		return err
	}

	var nURL = api + "?" + values.Encode()
	req, err := http.NewRequest(method, nURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rsp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err = json.NewDecoder(rsp.Body).Decode(result); err != nil {
		return err
	}
	return nil
}
