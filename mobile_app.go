package wxm

import "net/http"

type MobileApp struct {
	client *client
}

func NewMobileApp(appId, appSecret string) *MobileApp {
	var c = &MobileApp{}
	c.client = newClient(appId, appSecret)
	return c
}

func (m *MobileApp) With(opts ...Option) *MobileApp {
	var n = &MobileApp{}
	n.client = m.client.With(opts...)
	return n
}

func (m *MobileApp) SetAccessToken(accessToken string) {
	m.client.accessToken = accessToken
}

func (m *MobileApp) SetHTTPClient(client *http.Client) {
	m.client.client = client
}
