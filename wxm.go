package wxm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	kGetToken = "https://api.weixin.qq.com/cgi-bin/token"
)

type client struct {
	appId     string
	appSecret string
	Client    *http.Client

	mu    sync.Mutex
	token *Token
}

func newClient(appId, appSecret string) *client {
	var c = &client{}
	c.appId = appId
	c.appSecret = appSecret
	c.Client = http.DefaultClient
	return c
}

// GetToken 小程序、公众号-获取全局唯一后台接口调用凭据（access_token） https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *client) GetToken() (result string, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token != nil && c.token.AccessToken != "" && c.token.Valid() {
		return c.token.AccessToken, nil
	}
	c.token, err = c.getToken()
	if err != nil {
		return "", err
	}

	if c.token.Code != 0 {
		return "", errors.New(c.token.Msg)
	}

	return c.token.AccessToken, nil
}

func (c *client) RefreshToken() (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token, err = c.getToken()
	if err != nil {
		return err
	}
	if c.token.IsFailure() {
		return c.token.Error
	}
	return nil
}

func (c *client) getToken() (result *Token, err error) {
	var values = url.Values{}
	values.Add("appid", c.appId)
	values.Add("secret", c.appSecret)
	values.Add("grant_type", "client_credential")

	if err := c.requestWithoutAccessToken(http.MethodGet, kGetToken, nil, values, &result); err != nil {
		return nil, err
	}

	if result != nil {
		result.CreateTime = time.Now().Unix()
	}

	return result, nil
}

func (c *client) requestWithAccessToken(method, api string, param interface{}, values url.Values, result interface{}) error {
	var data, err = c.request(method, api, param, values, true, true)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}

func (c *client) requestWithoutAccessToken(method, api string, param interface{}, values url.Values, result interface{}) error {
	var data, err = c.request(method, api, param, values, false, false)
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}

func (c *client) request(method, api string, param interface{}, values url.Values, needAuth, retry bool) (result []byte, err error) {
	if values == nil {
		values = url.Values{}
	}

	if needAuth {
		accessToken, err := c.GetToken()
		if err != nil {
			return nil, err
		}
		values.Set("access_token", accessToken)
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
	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	result, err = io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if retry && string(result[11:16]) == strconv.Itoa(int(CodeInvalidCredential)) {
		if err = c.RefreshToken(); err != nil {
			return nil, err
		}
		return c.request(method, api, param, values, needAuth, false)
	}
	return result, nil
}

func (c *client) uploadWithRetry(method, api, fieldName, filePath string, values url.Values, needAuth bool, result interface{}) error {
	return c.upload(method, api, fieldName, filePath, values, needAuth, true, result)
}

func (c *client) upload(method, api, fieldName, filePath string, values url.Values, needAuth, retry bool, result interface{}) error {
	if values == nil {
		values = url.Values{}
	}

	if needAuth {
		accessToken, err := c.GetToken()
		if err != nil {
			return err
		}
		values.Set("access_token", accessToken)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldName, filePath)
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

	rsp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	if retry && string(bodyBytes[11:16]) == strconv.Itoa(int(CodeInvalidCredential)) {
		if err = c.RefreshToken(); err != nil {
			return err
		}
		return c.upload(method, api, fieldName, filePath, values, needAuth, false, result)
	}

	if err = json.Unmarshal(bodyBytes, result); err != nil {
		return err
	}
	return nil
}
