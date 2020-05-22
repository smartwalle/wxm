package wxm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const (
	kGetTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"

	kJSCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?grant_type=%s&appid=%s&secret=%s&js_code=%s"

	kGetUnLimitURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)

type Client struct {
	appId     string
	appSecret string
	client    *http.Client

	mu    sync.Mutex
	token *Token
}

func New(appId, appSecret string) *Client {
	var c = &Client{}
	c.appId = appId
	c.appSecret = appSecret
	c.client = http.DefaultClient
	return c
}

// GetToken 小程序、公众号-获取小程序全局唯一后台接口调用凭据（access_token） https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (this *Client) GetToken() (result string, err error) {
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

func (this *Client) RefreshToken() (err error) {
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

func (this *Client) getToken() (result *Token, err error) {
	var nURL = fmt.Sprintf(kGetTokenURL, "client_credential", this.appId, this.appSecret)
	req, err := http.NewRequest(http.MethodGet, nURL, nil)
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
	data, err := ioutil.ReadAll(rsp.Body)
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

func (this *Client) request(method, api string, param interface{}, values url.Values) (result []byte, err error) {
	return this.requestWithRetry(method, api, param, values, true)
}

func (this *Client) requestWithRetry(method, api string, param interface{}, values url.Values, retry bool) (result []byte, err error) {
	accessToken, err := this.GetToken()
	if err != nil {
		return nil, err
	}
	var nURL = fmt.Sprintf(api, accessToken)
	if values != nil {
		nURL = nURL + "&" + values.Encode()
	}

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
		return this.requestWithRetry(method, api, param, values, false)
	}
	return result, nil
}

// JSCode2Session 小程序-登录 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (this *Client) JSCode2Session(code string) (result *JSCode2SessionRsp, err error) {
	var nURL = fmt.Sprintf(kJSCode2SessionURL, "authorization_code", this.appId, this.appSecret, code)

	req, err := http.NewRequest(http.MethodGet, nURL, nil)
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
	data, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUnLimited 小程序-获取小程序码 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (this *Client) GetUnLimited(param GetUnlimitedParam) (result *GetUnlimitedRsp, err error) {
	data, err := this.request(http.MethodPost, kGetUnLimitURL, param, nil)
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if err = json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
		return result, nil
	}

	result = &GetUnlimitedRsp{}
	result.ErrCode = 0
	result.ErrMsg = "ok"
	result.Data = data

	return result, nil
}
