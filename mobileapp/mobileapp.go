package mobileapp

import "github.com/smartwalle/wxm"

type MobileApp struct {
	*wxm.Client
}

func New() *MobileApp {
	return &MobileApp{
		Client: wxm.New(),
	}
}
