package wxm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	kGetTokenURL = "https://api.weixin.qq.com/cgi-bin/token"
)

type client struct {
	appId     string
	appSecret string
	client    *http.Client

	mu    sync.Mutex
	token *Token
}

func newClient(appId, appSecret string) *client {
	var c = &client{}
	c.appId = appId
	c.appSecret = appSecret
	c.client = http.DefaultClient
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

	if this.token.ErrCode != 0 {
		return "", errors.New(this.token.ErrMsg)
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

	if this.token.ErrCode != 0 {
		return errors.New(this.token.ErrMsg)
	}

	return nil
}

func (this *client) getToken() (result *Token, err error) {
	var v = url.Values{}
	v.Add("appid", this.appId)
	v.Add("secret", this.appSecret)
	v.Add("grant_type", "client_credential")

	data, err := this.RequestWithoutAccessToken(http.MethodGet, kGetTokenURL, nil, v)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	if result != nil {
		result.CreateTime = time.Now().Unix()
	}

	return result, nil
}

func (this *client) RequestWithAccessToken(method, api string, param interface{}, values url.Values) (result []byte, err error) {
	return this.request(method, api, param, values, true, true)
}

func (this *client) RequestWithoutAccessToken(method, api string, param interface{}, values url.Values) (result []byte, err error) {
	return this.request(method, api, param, values, false, false)
}

func (this *client) request(method, api string, param interface{}, values url.Values, withAccessToken, retry bool) (result []byte, err error) {
	if values == nil {
		values = url.Values{}
	}

	if withAccessToken {
		accessToken, err := this.GetToken()
		if err != nil {
			return nil, err
		}

		values.Del("access_token")
		values.Add("access_token", accessToken)
	}

	var nURL = api + "?" + values.Encode()

	var body io.Reader
	if param != nil {
		data, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, nURL, body)
	if err != nil {
		return nil, err
	}
	rsp, err := this.client.Do(req)
	if rsp != nil && rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	result, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if retry && string(result[11:16]) == strconv.Itoa(int(CodeInvalidCredential)) {
		if err = this.RefreshToken(); err != nil {
			return nil, err
		}
		return this.request(method, api, param, values, withAccessToken, false)
	}
	return result, nil
}

func (this *client) Upload(method, api, fieldName, filePath string, values url.Values, withAccessToken bool) (result []byte, err error) {
	return this.upload(method, api, fieldName, filePath, values, withAccessToken, true)
}

func (this *client) upload(method, api, fieldName, filePath string, values url.Values, withAccessToken, retry bool) (result []byte, err error) {
	if values == nil {
		values = url.Values{}
	}

	if withAccessToken {
		accessToken, err := this.GetToken()
		if err != nil {
			return nil, err
		}

		values.Del("access_token")
		values.Add("access_token", accessToken)
	}

	var nURL = api + "?" + values.Encode()

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldName, filePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, nURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rsp, err := this.client.Do(req)
	if rsp != nil && rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	result, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if retry && string(result[11:16]) == strconv.Itoa(int(CodeInvalidCredential)) {
		if err = this.RefreshToken(); err != nil {
			return nil, err
		}
		return this.upload(method, api, fieldName, filePath, values, withAccessToken, false)
	}
	return result, nil
}
