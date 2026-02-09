package channel

import "github.com/smartwalle/wxm"

// GetFinderOverallRequest 获取电商概览数据请求参数
type GetFinderOverallRequest struct {
	Date string `json:"ds"` // 日期，格式YYYYMMDD
}

// GetFinderOverallResponse 获取电商概览数据响应
type GetFinderOverallResponse struct {
	wxm.Error
	Data *FinderOverallData `json:"data,omitempty"`
}

// FinderOverallData 电商概览数据
type FinderOverallData struct {
	PayGmv        int64 `json:"pay_gmv"`         // 成交金额，单位分
	LivePayGmv    int64 `json:"live_pay_gmv"`    // 直播成交金额，单位分
	FeedPayGmv    int64 `json:"feed_pay_gmv"`    // 短视频成交金额，单位分
	WindowPayGmv  int64 `json:"window_pay_gmv"`  // 橱窗成交金额，单位分
	ProductPayGmv int64 `json:"product_pay_gmv"` // 商品分享支付金额，单位分
	OtherPayGmv   int64 `json:"other_pay_gmv"`   // 其他渠道成交金额，单位分
}
