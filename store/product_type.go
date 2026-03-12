package store

import "github.com/smartwalle/wxm"

// ProductStatus 商品状态
type ProductStatus int

const (
	ProductStatusInitial      ProductStatus = 0  // 初始值
	ProductStatusOnShelf      ProductStatus = 5  // 上架
	ProductStatusRecycle      ProductStatus = 6  // 回收站
	ProductStatusDeleted      ProductStatus = 9  // 彻底删除，商品无法再进行任何操作
	ProductStatusOffShelf     ProductStatus = 11 // 自主下架
	ProductStatusViolationOff ProductStatus = 13 // 违规下架/风控系统下架
	ProductStatusDepositOff   ProductStatus = 14 // 保证金不足下架
	ProductStatusBrandExpired ProductStatus = 15 // 品牌过期下架
	ProductStatusBanned       ProductStatus = 20 // 商品被封禁
)

// GetProductListRequest 获取商品列表请求
type GetProductListRequest struct {
	Status   ProductStatus `json:"status,omitempty"`   // 商品状态，不填默认拉全部商品（不包含从未上架的草稿和回收站商品）
	PageSize int           `json:"page_size"`          // 每页数量（默认10，不超过30）
	NextKey  string        `json:"next_key,omitempty"` // 由上次请求返回，记录翻页的上下文
}

// GetProductListResponse 获取商品列表响应
type GetProductListResponse struct {
	wxm.Error
	ProductIDs []int64 `json:"product_ids"` // 商品id列表
	NextKey    string  `json:"next_key"`    // 本次翻页的上下文，用于请求下一页
	TotalNum   int     `json:"total_num"`   // 商品总数
}
