package wxm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	kAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"

	kJSCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?grant_type=%s&appid=%s&secret=%s&js_code=%s"

	kGetUnlimitURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)

type Client struct {
	appId     string
	appSecret string
	client    *http.Client

	mu          sync.Mutex
	accessToken *AccessToken
}

func New(appId, appSecret string) *Client {
	var c = &Client{}
	c.appId = appId
	c.appSecret = appSecret
	c.client = http.DefaultClient
	return c
}

func (this *Client) GetAccessToken() (result string, err error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if this.accessToken != nil && this.accessToken.AccessToken != "" && this.accessToken.Valid() {
		return this.accessToken.AccessToken, nil
	}
	this.accessToken, err = this.getAccessToken()
	if err != nil {
		return "", err
	}

	if this.accessToken.ErrCode != 0 {
		return "", errors.New(this.accessToken.ErrMsg)
	}

	return this.accessToken.AccessToken, nil
}

func (this *Client) RefreshAccessToken() (err error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.accessToken, err = this.getAccessToken()
	if err != nil {
		return err
	}

	if this.accessToken.ErrCode != 0 {
		return errors.New(this.accessToken.ErrMsg)
	}

	return nil
}

func (this *Client) getAccessToken() (result *AccessToken, err error) {
	var url = fmt.Sprintf(kAccessTokenURL, "client_credential", this.appId, this.appSecret)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

func (this *Client) request(method, api string, param interface{}) (result []byte, err error) {
	return this.request2(method, api, param, true)
}

func (this *Client) request2(method, api string, param interface{}, reTry bool) (result []byte, err error) {
	accessToken, err := this.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf(api, accessToken)

	var body io.Reader
	if param != nil {
		data, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, url, body)
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

	if reTry && string(result[11:16]) == strconv.Itoa(int(CodeInvalidCredential)) {
		if err = this.RefreshAccessToken(); err != nil {
			return nil, err
		}
		return this.request2(method, api, param, false)
	}
	return result, nil
}

// JSCode2Session 获取 session
func (this *Client) JSCode2Session(code string) (result *JSCode2SessionRsp, err error) {
	var url = fmt.Sprintf(kJSCode2SessionURL, "authorization_code", this.appId, this.appSecret, code)

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

// GetUnlimited 获取小程序码
func (this *Client) GetUnlimited(param GetUnlimitedParam) (result *GetUnlimitedRsp, err error) {
	data, err := this.request(http.MethodPost, kGetUnlimitURL, param)
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
