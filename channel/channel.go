package channel

import "github.com/smartwalle/wxm"

// Channel 视频号
type Channel struct {
	*wxm.Client
}

func New() *Channel {
	return &Channel{
		Client: wxm.New(),
	}
}
