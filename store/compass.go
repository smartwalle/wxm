package store

import (
	"context"
)

const (
	APIGetShopOverall     = "/channels/ec/compass/shop/overall/get"
	APIGetShopProductData = "/channels/ec/compass/shop/product/data/get"
)

// GetShopOverall 获取电商数据概览
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopoverall.html
func (s *Store) GetShopOverall(ctx context.Context, accessToken string, request GetShopOverallRequest) (response *GetShopOverallResponse, err error) {
	if err = s.Post(ctx, APIGetShopOverall, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetShopProductData 获取商品详细信息
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopproductdata.html
func (s *Store) GetShopProductData(ctx context.Context, accessToken string, request GetShopProductDataRequest) (response *GetShopProductDataResponse, err error) {
	if err = s.Post(ctx, APIGetShopProductData, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
