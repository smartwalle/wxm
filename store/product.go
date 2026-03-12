package store

import (
	"context"
)

const (
	APIGetProductList = "/channels/ec/product/list/get"
)

// GetProductList 获取商品列表
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-product/shop/api_getproductlist.html
func (s *Store) GetProductList(ctx context.Context, accessToken string, request GetProductListRequest) (response *GetProductListResponse, err error) {
	if err = s.Post(ctx, APIGetProductList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
