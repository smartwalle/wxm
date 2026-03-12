package store

import (
	"context"

	"github.com/smartwalle/wxm"
)

const (
	APIGetBasicInfo = "/channels/ec/basics/info/get"
)

// Store 微信小店
type Store struct {
	*wxm.Client
}

func New() *Store {
	return &Store{
		Client: wxm.New(),
	}
}

// GetBasicInfo 获取店铺基本信息
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_mmecapi_basicinfo.html
func (s *Store) GetBasicInfo(ctx context.Context, accessToken string) (response *GetBasicInfoResponse, err error) {
	if err = s.Get(ctx, APIGetBasicInfo, accessToken, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
