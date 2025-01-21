package wxm

type Website struct {
	client *client
}

func NewWebsite(appId, appSecret string, opts ...Option) *Website {
	var c = &Website{}
	c.client = newClient(appId, appSecret, opts...)
	return c
}
