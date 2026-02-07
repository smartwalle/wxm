package officialaccount

import (
	"github.com/smartwalle/wxm"
)

type OfficialAccount struct {
	*wxm.Client
}

func New() *OfficialAccount {
	return &OfficialAccount{
		Client: wxm.New(),
	}
}
