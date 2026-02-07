package website

import "github.com/smartwalle/wxm"

type Website struct {
	*wxm.Client
}

func New() *Website {
	return &Website{
		Client: wxm.New(),
	}
}
