package wxm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	kAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"

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

// GetUnlimited 获取小程序码
func (this *Client) GetUnlimited(param GetUnlimitedParam) (result *GetUnlimitedRsp, err error) {
	accessToken, err := this.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf(kGetUnlimitURL, accessToken)

	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
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

	data, err = ioutil.ReadAll(rsp.Body)
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
