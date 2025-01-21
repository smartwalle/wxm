package wxm

type MobileApp struct {
	client *client
}

func NewMobileApp(appId, appSecret string, opts ...Option) *MobileApp {
	var c = &MobileApp{}
	c.client = newClient(appId, appSecret, opts...)
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
