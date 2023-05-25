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
func (this *client) GetToken() (result string, err error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if this.token != nil && this.token.AccessToken != "" && this.token.Valid() {
		return this.token.AccessToken, nil
	}
	this.token, err = this.getToken()
	if err != nil {
		return "", err
	}

	if this.token.Code != 0 {
		return "", errors.New(this.token.Msg)
	}

	return this.token.AccessToken, nil
}

func (this *client) RefreshToken() (err error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.token, err = this.getToken()
	if err != nil {
		return err
	}
	if this.token.IsFailure() {
		return this.token.Error
	}
	return nil
}

func (this *client) getToken() (result *Token, err error) {
	var values = url.Values{}
	values.Add("appid", this.appId)
	values.Add("secret", this.appSecret)
	values.Add("grant_type", "client_credential")

	if err := this.requestWithoutAccessToken(http.MethodGet, kGetToken, nil, values, &result); err != nil {
		return nil, err
	}

	if result != nil {
		result.CreateTime = time.Now().Unix()
	}

	return result, nil
}

func (this *client) requestWithAccessToken(method, api string, param interface{}, values url.Values, result interface{}) error {
	var data, err = this.request(method, api, param, values, true, true)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}

func (this *client) requestWithoutAccessToken(method, api string, param interface{}, values url.Values, result interface{}) error {
	var data, err = this.request(method, api, param, values, false, false)
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}

func (this *client) request(method, api string, param interface{}, values url.Values, needAuth, retry bool) (result []byte, err error) {
	if values == nil {
		values = url.Values{}
	}

	if needAuth {
		accessToken, err := this.GetToken()
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
	rsp, err := this.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	result, err = io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if retry && string(result[11:16]) == strconv.Itoa(int(CodeInvalidCredential)) {
		if err = this.RefreshToken(); err != nil {
			return nil, err
		}
		return this.request(method, api, param, values, needAuth, false)
	}
	return result, nil
}

func (this *client) upload(method, api, fieldName, filePath string, values url.Values, needAuth bool, result interface{}) error {
	return this._upload(method, api, fieldName, filePath, values, needAuth, true, result)
}

func (this *client) _upload(method, api, fieldName, filePath string, values url.Values, needAuth, retry bool, result interface{}) error {
	if values == nil {
		values = url.Values{}
	}

	if needAuth {
		accessToken, err := this.GetToken()
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

	rsp, err := this.Client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	if retry && string(bodyBytes[11:16]) == strconv.Itoa(int(CodeInvalidCredential)) {
		if err = this.RefreshToken(); err != nil {
			return err
		}
		return this._upload(method, api, fieldName, filePath, values, needAuth, false, result)
	}

	if err = json.Unmarshal(bodyBytes, result); err != nil {
		return err
	}
	return nil
}
