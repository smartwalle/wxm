package store

import "github.com/smartwalle/wxm"

// OrderStatus 订单状态
type OrderStatus int

const (
	OrderStatusToPay            OrderStatus = 10  // 待付款
	OrderStatusGiftToReceive    OrderStatus = 12  // 礼物待收下
	OrderStatusGroupPending     OrderStatus = 13  // 一起买待成团
	OrderStatusToShip           OrderStatus = 20  // 待发货（包括部分发货）
	OrderStatusPartiallyShipped OrderStatus = 21  // 部分发货
	OrderStatusToReceive        OrderStatus = 30  // 待收货（包括部分发货）
	OrderStatusCompleted        OrderStatus = 100 // 完成
	OrderStatusCancelled        OrderStatus = 250 // 订单取消（包括未付款取消，售后取消等）
)

// TimeRange 时间范围
type TimeRange struct {
	StartTime int64 `json:"start_time"` // 秒级时间戳（距离end_time不可超过7天）
	EndTime   int64 `json:"end_time"`   // 秒级时间戳（距离start_time不可超过7天）
}

// GetOrderListRequest 获取订单列表请求参数
type GetOrderListRequest struct {
	CreateTimeRange *TimeRange  `json:"create_time_range"`   // 订单创建时间范围，时间范围至少填一个
	UpdateTimeRange *TimeRange  `json:"update_time_range"`   // 订单更新时间范围，时间范围至少填一个
	Status          OrderStatus `json:"status,omitempty"`    // 订单状态
	OpenID          string      `json:"openid,omitempty"`    // 买家身份标识
	PageSize        int         `json:"page_size,omitempty"` // 每页数量(不超过100)
	NextKey         string      `json:"next_key,omitempty"`  // 分页参数
}

// GetOrderListResponse 获取订单列表响应
type GetOrderListResponse struct {
	wxm.Error
	OrderIDList []string `json:"order_id_list"` // 订单号列表
	NextKey     string   `json:"next_key"`      // 分页参数
	HasMore     bool     `json:"has_more"`      // 是否还有下一页
}
