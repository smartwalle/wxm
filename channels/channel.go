package channels

import "github.com/smartwalle/wxm"

// Channels 视频号
type Channels struct {
	*wxm.Client
}

func New() *Channels {
	return &Channels{
		Client: wxm.New(),
	}
}
