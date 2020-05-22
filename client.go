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

func (this *client) request(method, api string, param interface{}, values url.Values) (result []byte, err error) {
	return this.requestWithRetry(method, api, param, values, true)
}

func (this *client) requestWithRetry(method, api string, param interface{}, values url.Values, retry bool) (result []byte, err error) {
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
