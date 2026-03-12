package store

import (
	"context"
)

const (
	APIGetOrderList = "/channels/ec/order/list/get"
)

// GetOrderList 获取订单列表
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getorderlist.html
func (s *Store) GetOrderList(ctx context.Context, accessToken string, request GetOrderListRequest) (response *GetOrderListResponse, err error) {
	if err = s.Post(ctx, APIGetOrderList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
