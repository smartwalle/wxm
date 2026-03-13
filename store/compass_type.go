package store

import "github.com/smartwalle/wxm"

// GetShopOverallRequest 获取电商数据概览请求参数
type GetShopOverallRequest struct {
	Date string `json:"ds"` // 日期，格式YYYYMMDD
}

// GetShopOverallResponse 获取电商数据概览响应
type GetShopOverallResponse struct {
	wxm.Error
	Data *ShopOverallData `json:"data,omitempty"` // 电商数据
}

// ShopOverallData 电商数据
type ShopOverallData struct {
	PayGMV         string `json:"pay_gmv"`          // 成交金额，单位分
	PayUV          string `json:"pay_uv"`           // 成交人数
	PayOrderCnt    string `json:"pay_order_cnt"`    // 成交订单数
	PayRefundGMV   string `json:"pay_refund_gmv"`   // 成交退款金额，单位分
	ProductClickUV string `json:"product_click_uv"` // 商品点击人数
	LivePayGMV     string `json:"live_pay_gmv"`     // 直播成交金额，单位分
	FeedPayGMV     string `json:"feed_pay_gmv"`     // 短视频成交金额，单位分
}

// GetShopProductDataRequest 获取商品详细信息请求参数
type GetShopProductDataRequest struct {
	Date      string `json:"ds"`         // 日期，格式YYYYMMDD
	ProductID string `json:"product_id"` // 商品id
}

// GetShopProductDataResponse 获取商品详细信息响应
type GetShopProductDataResponse struct {
	wxm.Error
	ProductInfo *ShopProductInfo `json:"product_info,omitempty"` // 商品详细信息
}

// ShopProductInfo 商品详细信息
type ShopProductInfo struct {
	ProductID        string               `json:"product_id"`         // 商品id
	HeadImgURL       string               `json:"head_img_url"`       // 商品图
	Title            string               `json:"title"`              // 商品标题
	Price            string               `json:"price"`              // 商品价格，单位分
	FirstCategoryID  string               `json:"first_category_id"`  // 商品一级类目
	SecondCategoryID string               `json:"second_category_id"` // 商品二级类目
	ThirdCategoryID  string               `json:"third_category_id"`  // 商品三级类目
	Data             *ShopProductDataInfo `json:"data,omitempty"`     // 详细数据
}

// ShopProductDataInfo 详细数据
type ShopProductDataInfo struct {
	PayGMV                     string   `json:"pay_gmv"`                                // 成交金额，单位分
	CreateGMV                  string   `json:"create_gmv"`                             // 下单金额，单位分
	CreateCnt                  string   `json:"create_cnt"`                             // 下单订单数
	CreateUV                   string   `json:"create_uv"`                              // 下单人数
	CreateProductCnt           string   `json:"create_product_cnt"`                     // 下单件数
	PayCnt                     string   `json:"pay_cnt"`                                // 成交订单数
	PayUV                      string   `json:"pay_uv"`                                 // 成交人数
	PayProductCnt              string   `json:"pay_product_cnt"`                        // 成交件数
	PurePayGMV                 string   `json:"pure_pay_gmv"`                           // 成交金额（剔除退款）
	PayGMVPerUV                string   `json:"pay_gmv_per_uv"`                         // 成交客单价（剔除退款）
	SellerActualSettleAmount   string   `json:"seller_actual_settle_amount"`            // 实际结算金额，单位分
	PlatformActualCommission   string   `json:"platform_actual_commission"`             // 实际服务费金额，单位分
	FinderUinActualCommission  string   `json:"finderuin_actual_commission"`            // 实际达人佣金支出，单位分
	CaptainActualCommission    string   `json:"captain_actual_commission"`              // 实际团长佣金支出，单位分
	SellerPredictSettleAmount  string   `json:"seller_predict_settle_amount"`           // 预估结算金额，单位分
	PlatformPredictCommission  string   `json:"platform_predict_commission"`            // 预估服务费金额，单位分
	FinderUinPredictCommission string   `json:"finderuin_predict_commission"`           // 预估达人佣金支出，单位分
	CaptainPredictCommission   string   `json:"captain_predict_commission"`             // 预估团长佣金支出，单位分
	ProductClickUV             string   `json:"product_click_uv"`                       // 商品点击人数
	ProductClickCnt            string   `json:"product_click_cnt"`                      // 商品点击次数
	PayRefundGMV               string   `json:"pay_refund_gmv"`                         // 成交退款金额，单位分
	PayRefundUV                string   `json:"pay_refund_uv"`                          // 成交退款人数
	PayRefundRatio             *float64 `json:"pay_refund_ratio,omitempty"`             // 成交退款率
	PayRefundAfterSendRatio    *float64 `json:"pay_refund_after_send_ratio,omitempty"`  // 发货后成交退款率
	PayRefundCnt               string   `json:"pay_refund_cnt"`                         // 成交退款订单数
	PayRefundProductCnt        string   `json:"pay_refund_product_cnt"`                 // 成交退款件数
	PayRefundBeforeSendRatio   *float64 `json:"pay_refund_before_send_ratio,omitempty"` // 发货前成交退款率
	RefundGMV                  *string  `json:"refund_gmv,omitempty"`                   // 退款金额，单位分
	RefundProductCnt           string   `json:"refund_product_cnt"`                     // 退款件数
	RefundCnt                  string   `json:"refund_cnt"`                             // 退款订单数
	RefundUV                   string   `json:"refund_uv"`                              // 退款人数
}
