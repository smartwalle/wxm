package store

import (
	"context"
)

const (
	APIGetShopOverall                 = "/channels/ec/compass/shop/overall/get"
	APIGetShopProductData             = "/channels/ec/compass/shop/product/data/get"
	APIGetShopFinderAuthorizationList = "/channels/ec/compass/shop/finder/authorization/list/get"
	APIGetShopFinderList              = "/channels/ec/compass/shop/finder/list/get"
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

// GetShopFinderAuthorizationList 获取授权视频号列表
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderauthorizationlist.html
func (s *Store) GetShopFinderAuthorizationList(ctx context.Context, accessToken string) (response *GetShopFinderAuthorizationListResponse, err error) {
	var request = struct{}{}
	if err = s.Post(ctx, APIGetShopFinderAuthorizationList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetShopFinderList 获取带货达人列表
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderlist.html
func (s *Store) GetShopFinderList(ctx context.Context, accessToken string, request GetShopFinderListRequest) (response *GetShopFinderListResponse, err error) {
	if err = s.Post(ctx, APIGetShopFinderList, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
