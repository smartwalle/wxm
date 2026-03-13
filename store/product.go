package store

import (
	"context"
	"strconv"

	"github.com/smartwalle/wxm"
)

const (
	APIGetProductList = "/channels/ec/product/list/get"
	APIGetProduct     = "/channels/ec/product/get"
	APIGetStock       = "/channels/ec/product/stock/get"
	APIBatchGetStock  = "/channels/ec/product/stock/batchget"
)

// GetProductList 获取商品列表
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-product/shop/api_getproductlist.html
func (s *Store) GetProductList(ctx context.Context, accessToken string, request GetProductListRequest) (response *GetProductListResponse, err error) {
	var aux = struct {
		wxm.Error
		ProductIds []int64 `json:"product_ids"` // 商品id列表
		NextKey    string  `json:"next_key"`    // 本次翻页的上下文，用于请求下一页
		TotalNum   int     `json:"total_num"`   // 商品总数
	}{}

	if err = s.Post(ctx, APIGetProductList, accessToken, request, nil, &aux); err != nil {
		return nil, err
	}
	response = &GetProductListResponse{
		Error:      aux.Error,
		ProductIds: make([]string, 0, len(aux.ProductIds)),
		NextKey:    aux.NextKey,
		TotalNum:   aux.TotalNum,
	}

	for _, productId := range aux.ProductIds {
		response.ProductIds = append(response.ProductIds, strconv.FormatInt(productId, 10))
	}

	return response, nil
}

// GetProduct 获取商品
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-product/shop/api_getproduct.html
func (s *Store) GetProduct(ctx context.Context, accessToken string, request GetProductRequest) (response *GetProductResponse, err error) {
	if err = s.Post(ctx, APIGetProduct, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetStock 获取库存
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-product/stock/api_getstock.html
func (s *Store) GetStock(ctx context.Context, accessToken string, request GetStockRequest) (response *GetStockResponse, err error) {
	if err = s.Post(ctx, APIGetStock, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// BatchGetStock 批量获取库存信息
//
//	接口文档：https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-product/stock/api_batchgetstock.html
func (s *Store) BatchGetStock(ctx context.Context, accessToken string, request BatchGetStockRequest) (response *BatchGetStockResponse, err error) {
	if err = s.Post(ctx, APIBatchGetStock, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
