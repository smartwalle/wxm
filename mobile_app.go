package wxm

type MobileApp struct {
	client *client
}

func NewMobileApp(appId, appSecret string) *MobileApp {
	var c = &MobileApp{}
	c.client = newClient(appId, appSecret)
	return c
}
