package wxm

type Website struct {
	*client
}

func NewWebsite(appId, appSecret string) *Website {
	var c = &Website{}
	c.client = newClient(appId, appSecret)
	return c
}
