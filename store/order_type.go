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

// GetOrderRequest 获取订单详情请求
type GetOrderRequest struct {
	OrderID string `json:"order_id"` // 订单ID，可从获取订单列表中获得
}

// GetOrderResponse 获取订单详情响应
type GetOrderResponse struct {
	wxm.Error
	Order *Order `json:"order,omitempty"` // 订单结构
}

// Order 订单结构
type Order struct {
	OrderID             string              `json:"order_id"`                        // 订单ID
	CreateTime          int64               `json:"create_time"`                     // 创建时间，秒级时间戳
	UpdateTime          int64               `json:"update_time"`                     // 更新时间，秒级时间戳
	Status              OrderStatus         `json:"status"`                          // 订单状态
	OrderDetail         *OrderDetail        `json:"order_detail,omitempty"`          // 订单详细数据信息
	AftersaleDetail     *AftersaleDetail    `json:"aftersale_detail,omitempty"`      // 售后信息
	OpenID              string              `json:"openid,omitempty"`                // 订单归属人身份标识
	UnionID             string              `json:"unionid,omitempty"`               // 订单归属人在开放平台的唯一标识符
	IsPresent           bool                `json:"is_present,omitempty"`            // 是否礼物订单
	PresentNote         string              `json:"present_note,omitempty"`          // 礼物订单留言
	PresentGiverOpenID  string              `json:"present_giver_openid,omitempty"`  // 礼物订单赠送者openid
	PresentGiverUnionID string              `json:"present_giver_unionid,omitempty"` // 礼物订单赠送者在开放平台的唯一标识符
	PresentOrderIDStr   string              `json:"present_order_id_str,omitempty"`  // 礼物订单ID
	PresentSendType     int                 `json:"present_send_type,omitempty"`     // 礼物单类型
	OrderPresentInfo    *OrderPresentInfo   `json:"order_present_info,omitempty"`    // 订单对应礼物单信息
	IsFlashSaleOrder    bool                `json:"is_flash_sale_order,omitempty"`   // 是否闪购订单
	IntraCityOrderInfo  *IntraCityOrderInfo `json:"intra_city_order_info,omitempty"` // 同城单信息
}

// OrderDetail 订单详细数据信息
type OrderDetail struct {
	ProductInfos     []OrderProduct    `json:"product_infos"`                // 商品列表
	PayInfo          *PayInfo          `json:"pay_info,omitempty"`           // 支付信息
	PriceInfo        *PriceInfo        `json:"price_info,omitempty"`         // 价格信息
	DeliveryInfo     *DeliveryInfo     `json:"delivery_info,omitempty"`      // 配送信息
	ExtInfo          *ExtInfo          `json:"ext_info,omitempty"`           // 额外信息
	CouponInfo       *CouponInfo       `json:"coupon_info,omitempty"`        // 优惠券信息
	CommissionInfos  []CommissionInfo  `json:"commission_infos,omitempty"`   // 分佣信息
	SharerInfo       *SharerInfo       `json:"sharer_info,omitempty"`        // 分享员信息【已经下线，不再维护】
	SettleInfo       *SettleInfo       `json:"settle_info,omitempty"`        // 结算信息
	SKUSharerInfos   []SKUSharerInfo   `json:"sku_sharer_infos,omitempty"`   // 分享员信息【已经下线，不再维护】
	AgentInfo        *AgentInfo        `json:"agent_info,omitempty"`         // 授权账号信息
	SourceInfos      []SourceInfo      `json:"source_infos,omitempty"`       // 订单来源信息
	RefundInfo       *RefundInfo       `json:"refund_info,omitempty"`        // 订单退款信息
	GreetingCardInfo *GreetingCardInfo `json:"greeting_card_info,omitempty"` // 需代写的商品贺卡信息
	CustomInfo       *CustomInfo       `json:"custom_info,omitempty"`        // 商品定制信息
}

// OrderProduct 商品信息
type OrderProduct struct {
	ProductID                               string                   `json:"product_id"`                                  // 商品id
	SKUID                                   string                   `json:"sku_id"`                                      // 商品skuid
	ThumbImg                                string                   `json:"thumb_img"`                                   // sku小图
	SalePrice                               int64                    `json:"sale_price"`                                  // 售卖单价，单位为分
	SKUCnt                                  int                      `json:"sku_cnt"`                                     // sku数量
	Title                                   string                   `json:"title"`                                       // 商品标题
	OnAftersaleSKUCnt                       int                      `json:"on_aftersale_sku_cnt"`                        // 正在售后/退款流程中的 sku 数量
	FinishAftersaleSKUCnt                   int                      `json:"finish_aftersale_sku_cnt"`                    // 完成售后/退款的 sku 数量
	SKUCode                                 string                   `json:"sku_code"`                                    // sku编码（商家自定义编码）
	MarketPrice                             int64                    `json:"market_price"`                                // 市场单价，单位为分
	SKUAttrs                                []Attribute              `json:"sku_attrs"`                                   // sku属性
	RealPrice                               int64                    `json:"real_price"`                                  // sku实付总价
	OutProductID                            string                   `json:"out_product_id"`                              // 商品外部spuid
	OutSKUID                                string                   `json:"out_sku_id"`                                  // 商品外部skuid
	IsDiscounted                            bool                     `json:"is_discounted"`                               // 是否有商家优惠金额
	EstimatePrice                           int64                    `json:"estimate_price"`                              // 使用所有优惠后sku总价
	IsChangePrice                           bool                     `json:"is_change_price"`                             // 是否修改过价格
	ChangePrice                             int64                    `json:"change_price"`                                // 改价后sku总价
	OutWarehouseID                          string                   `json:"out_warehouse_id"`                            // 区域库存id
	SKUDeliverInfo                          *SKUDeliverInfo          `json:"sku_deliver_info"`                            // 商品发货信息
	ExtraService                            *ProductExtraService     `json:"extra_service"`                               // 商品额外服务信息
	UseDeduction                            bool                     `json:"use_deduction"`                               // 是否使用了会员积分抵扣
	DeductionPrice                          int64                    `json:"deduction_price"`                             // 会员积分抵扣金额，单位为分
	OrderProductCouponInfoList              []OrderProductCouponInfo `json:"order_product_coupon_info_list"`              // 商品优惠券信息
	DeliveryDeadline                        int64                    `json:"delivery_deadline"`                           // 商品发货时效
	MerchantDiscountedPrice                 int64                    `json:"merchant_discounted_price"`                   // 商家优惠金额，单位为分
	FinderDiscountedPrice                   int64                    `json:"finder_discounted_price"`                     // 达人优惠金额，单位为分
	IsFreeGift                              int                      `json:"is_free_gift"`                                // 是否赠品
	VIPDiscountedPrice                      int64                    `json:"vip_discounted_price"`                        // 订单内商品维度会员权益优惠金额
	ProductUniqueID                         string                   `json:"product_unique_id"`                           // 商品常量编号
	ChangeSKUInfo                           *ChangeSKUInfo           `json:"change_sku_info"`                             // 更换sku信息
	FreeGiftInfo                            *FreeGiftInfo            `json:"free_gift_info"`                              // 赠品信息
	BulkbuyDiscountedPrice                  int64                    `json:"bulkbuy_discounted_price"`                    // 订单内商品维度一起买优惠金额
	NationalSubsidyDiscountedPrice          int64                    `json:"national_subsidy_discounted_price"`           // 订单内商品维度国补优惠金额
	DropshipInfo                            *DropshipInfo            `json:"dropship_info"`                               // 代发相关信息
	IsFlashSale                             bool                     `json:"is_flash_sale"`                               // 是否闪购商品
	NationalSubsidyMerchantDiscountedPrice  int64                    `json:"national_subsidy_merchant_discounted_price"`  // 地方补贴优惠金额(商家出资)
	PlatformActivityMerchantDiscountedPrice int64                    `json:"platform_activity_merchant_discounted_price"` // 活动商家补贴
	CashCouponDiscountedPrice               int64                    `json:"cash_coupon_discounted_price"`                // 订单内商品维度平台券优惠金额
	LimitedDiscountDiscountedPrice          int64                    `json:"limited_discount_discounted_price"`           // 限时抢购优惠金额
}

// ProductExtraService 商品额外服务信息
type ProductExtraService struct {
	SevenDayReturn   int `json:"seven_day_return"`  // 7天无理由：0：不支持，1：支持
	FreightInsurance int `json:"freight_insurance"` // 商家运费险：0：不支持，1：支持
}

// OrderProductCouponInfo 商品优惠券信息
type OrderProductCouponInfo struct {
	UserCouponID    string `json:"user_coupon_id"`   // 用户优惠券id
	CouponType      int    `json:"coupon_type"`      // 优惠券类型
	DiscountedPrice int64  `json:"discounted_price"` // 优惠金额，单位为分
	CouponID        string `json:"coupon_id"`        // 优惠券id
}

// ChangeSKUInfo 更换sku信息
type ChangeSKUInfo struct {
	PreshipmentChangeSKUState int   `json:"preshipment_change_sku_state"` // 发货前更换sku状态
	OldSKUID                  int64 `json:"old_sku_id"`                   // 原sku_id
	NewSKUID                  int64 `json:"new_sku_id"`                   // 用户申请更换的sku_id
	DDLTimeStamp              int64 `json:"ddl_time_stamp"`               // 商家处理请求的最后时间
}

// FreeGiftInfo 赠品信息
type FreeGiftInfo struct {
	MainProductList []MainProductInfo `json:"main_product_list"` // 赠品对应的主品信息
}

// MainProductInfo 赠品对应的主品信息
type MainProductInfo struct {
	GiftCnt   int    `json:"gift_cnt"`   // 赠品数量
	TaskID    int    `json:"task_id"`    // 活动id
	ProductID string `json:"product_id"` // 商品id
	SKUID     int    `json:"sku_id"`     // 主品sku_id
}

// DropshipInfo 代发相关信息
type DropshipInfo struct {
	DSOrderID int `json:"ds_order_id"` // 代发单号
}

// PayInfo 支付信息
type PayInfo struct {
	PaymentMethod int    `json:"payment_method"` // 支付方式
	PayTime       int64  `json:"pay_time"`       // 支付时间，秒级时间戳
	TransactionID string `json:"transaction_id"` // 支付订单号
	PrepayID      string `json:"prepay_id"`      // 预支付ID
	PrepayTime    int64  `json:"prepay_time"`    // 预支付时间
}

// PriceInfo 价格信息
type PriceInfo struct {
	ProductPrice                            int64 `json:"product_price"`                               // 商品总价，单位为分
	OrderPrice                              int64 `json:"order_price"`                                 // 用户实付金额，单位为分
	Freight                                 int64 `json:"freight"`                                     // 运费，单位为分
	DiscountedPrice                         int64 `json:"discounted_price"`                            // 商家优惠金额，单位为分
	IsDiscounted                            bool  `json:"is_discounted"`                               // 是否有商家优惠券优惠
	OriginalOrderPrice                      int64 `json:"original_order_price"`                        // 订单原始价格，单位为分
	EstimateProductPrice                    int64 `json:"estimate_product_price"`                      // 商品预估价格，单位为分
	ChangeDownPrice                         int64 `json:"change_down_price"`                           // 改价后降低金额，单位为分
	ChangeFreight                           int64 `json:"change_freight"`                              // 改价后运费，单位为分
	IsChangeFreight                         bool  `json:"is_change_freight"`                           // 是否修改运费
	UseDeduction                            bool  `json:"use_deduction"`                               // 是否使用了会员积分抵扣
	DeductionPrice                          int64 `json:"deduction_price"`                             // 会员积分抵扣金额，单位为分
	MerchantReceievePrice                   int64 `json:"merchant_receieve_price"`                     // 商家实收金额，单位为分
	MerchantDiscountedPrice                 int64 `json:"merchant_discounted_price"`                   // 商家优惠金额，单位为分
	FinderDiscountedPrice                   int64 `json:"finder_discounted_price"`                     // 达人优惠金额，单位为分
	VIPDiscountedPrice                      int64 `json:"vip_discounted_price"`                        // 订单维度会员权益优惠金额
	BulkbuyDiscountedPrice                  int64 `json:"bulkbuy_discounted_price"`                    // 订单维度一起买优惠金额
	NationalSubsidyDiscountedPrice          int64 `json:"national_subsidy_discounted_price"`           // 订单维度国补优惠金额
	CashCouponDiscountedPrice               int64 `json:"cash_coupon_discounted_price"`                // 订单维度平台券优惠金额
	NationalSubsidyMerchantDiscountedPrice  int64 `json:"national_subsidy_merchant_discounted_price"`  // 订单维度地方补贴优惠(商家出资)
	PlatformActivityMerchantDiscountedPrice int64 `json:"platform_activity_merchant_discounted_price"` // 活动商家补贴
	LimitedDiscountDiscountedPrice          int64 `json:"limited_discount_discounted_price"`           // 限时抢购优惠金额
}

// DeliveryInfo 配送信息
type DeliveryInfo struct {
	AddressInfo         *AddressInfo          `json:"address_info"`          // 地址信息
	DeliveryProductInfo []DeliveryProductInfo `json:"delivery_product_info"` // 发货物流信息
	ShipDoneTime        int64                 `json:"ship_done_time"`        // 发货完成时间，秒级时间戳
	DeliverMethod       int                   `json:"deliver_method"`        // 订单发货方式
	AddressUnderReview  *AddressInfo          `json:"address_under_review"`  // 用户下单后申请修改收货地址
	AddressApplyTime    int64                 `json:"address_apply_time"`    // 修改地址申请时间，秒级时间戳
	RechargeInfo        *RechargeInfo         `json:"recharge_info"`         // 虚拟商品充值账户信息
	EwaybillOrderCode   string                `json:"ewaybill_order_code"`   // 电子面单跨店铺取号的订单密钥
	QualityInspectType  int                   `json:"quality_inspect_type"`  // 订单质检类型
	QualityInspectInfo  *QualityInspectInfo   `json:"quality_inspect_info"`  // 质检信息
	DropshipFlag        int                   `json:"dropship_flag"`         // 供货商代发标记
	PredictDeliveryTime int64                 `json:"predict_delivery_time"` // 预计发货时间
	DeliveryTimeType    int                   `json:"delivery_time_type"`    // 发货时效类型
}

// AddressInfo 地址信息
type AddressInfo struct {
	UserName              string            `json:"user_name"`                // 收货人姓名
	PostalCode            string            `json:"postal_code"`              // 邮编
	ProvinceName          string            `json:"province_name"`            // 省份
	CityName              string            `json:"city_name"`                // 城市
	CountyName            string            `json:"county_name"`              // 区
	DetailInfo            string            `json:"detail_info"`              // 详细地址
	TelNumber             string            `json:"tel_number"`               // 联系方式
	HouseNumber           string            `json:"house_number"`             // 门牌号码
	VirtualOrderTelNumber string            `json:"virtual_order_tel_number"` // 虚拟发货订单联系方式
	UseTelNumber          int               `json:"use_tel_number"`           // 是否使用虚拟号码
	TelNumberExtInfo      *TelNumberExtInfo `json:"tel_number_ext_info"`      // 额外的联系方式信息
	HashCode              string            `json:"hash_code"`                // 标识当前店铺下一个唯一的用户收货地址
}

// TelNumberExtInfo 额外的联系方式信息
type TelNumberExtInfo struct {
	RealTelNumber        string `json:"real_tel_number"`         // 脱敏手机号
	VirtualTelNumber     string `json:"virtual_tel_number"`      // 完整的虚拟号码
	VirtualTelExpireTime int64  `json:"virtual_tel_expire_time"` // 虚拟号过期时间
	HasDelayTimes        int    `json:"has_delay_times"`         // 已延期次数
	GetVirtualTelCnt     int    `json:"get_virtual_tel_cnt"`     // 主动兑换虚拟号码次数
}

// DeliveryProductInfo 发货物流信息
type DeliveryProductInfo struct {
	WaybillID    string               `json:"waybill_id"`    // 快递单号
	DeliveryID   string               `json:"delivery_id"`   // 快递公司编码
	DeliveryName string               `json:"delivery_name"` // 快递公司名称
	DeliveryTime int64                `json:"delivery_time"` // 发货时间，秒级时间戳
	DeliverType  int                  `json:"deliver_type"`  // 配送方式
	ProductInfos []PackageProductInfo `json:"product_infos"` // 包裹中商品信息
}

// PackageProductInfo 包裹中商品信息
type PackageProductInfo struct {
	ProductID  int `json:"product_id"`  // 商品id
	SKUID      int `json:"sku_id"`      // sku_id
	ProductCnt int `json:"product_cnt"` // 商品数量
}

// RechargeInfo 虚拟商品充值账户信息
type RechargeInfo struct {
	AccountNo   string `json:"account_no"`   // 虚拟商品充值账号
	AccountType string `json:"account_type"` // 账号充值类型
	WXOpenID    string `json:"wx_openid"`    // 微信openid
}

// QualityInspectInfo 质检信息
type QualityInspectInfo struct {
	InspectStatus int `json:"inspect_status"` // 质检状态
}

// ExtInfo 额外信息
type ExtInfo struct {
	CustomerNotes              string `json:"customer_notes"`               // 用户备注
	MerchantNotes              string `json:"merchant_notes"`               // 商家备注
	ConfirmReceiptTime         int64  `json:"confirm_receipt_time"`         // 确认收货时间
	FinderID                   string `json:"finder_id"`                    // 视频号id
	LiveID                     string `json:"live_id"`                      // 直播id
	OrderScene                 int    `json:"order_scene"`                  // 下单场景
	VIPOrderSessionID          string `json:"vip_order_session_id"`         // 会员权益-session_id
	CommissionHandlingProgress int    `json:"commission_handling_progress"` // 用于判断分佣单是否已生成
}

// CouponInfo 优惠券信息
type CouponInfo struct {
	UserCouponID string `json:"user_coupon_id"` // 用户优惠券id
}

// CommissionInfo 分佣信息
type CommissionInfo struct {
	SKUID        string `json:"sku_id"`       // 商品skuid
	Nickname     string `json:"nickname"`     // 分账方昵称
	Type         int    `json:"type"`         // 分账方类型
	Status       int    `json:"status"`       // 分账状态
	Amount       int64  `json:"amount"`       // 分账金额
	FinderID     string `json:"finder_id"`    // 达人视频号id
	Openfinderid string `json:"openfinderid"` // 达人openfinderid
	TalentID     string `json:"talent_id"`    // 新带货达人 id
	AgencyID     string `json:"agency_id"`    // 带货机构 id
}

// SharerInfo 分享员信息【已经下线，不再维护】
type SharerInfo struct {
	SharerOpenID     string `json:"sharer_openid"`     // 分享员openid
	SharerUnionID    string `json:"sharer_unionid"`    // 分享员unionid
	SharerType       int    `json:"sharer_type"`       // 分享员类型
	ShareScene       int    `json:"share_scene"`       // 分享场景
	HandlingProgress int    `json:"handling_progress"` // 分享员数据是否已经解析完成
}

// SettleInfo 结算信息
type SettleInfo struct {
	PredictCommissionFee    int64 `json:"predict_commission_fee"`    // 预计技术服务费，单位为分
	CommissionFee           int64 `json:"commission_fee"`            // 实际技术服务费，单位为分
	PredictWecoinCommission int64 `json:"predict_wecoin_commission"` // 预计人气卡返佣金额，单位为分
	WecoinCommission        int64 `json:"wecoin_commission"`         // 实际人气卡返佣金额，单位为分
	SettleTime              int64 `json:"settle_time"`               // 商家结算时间
}

// SKUSharerInfo 分享员信息【已经下线，不再维护】
type SKUSharerInfo struct {
	SharerOpenID  string `json:"sharer_openid"`  // 分享员openid
	SharerUnionID string `json:"sharer_unionid"` // 分享员unionid
	SharerType    int    `json:"sharer_type"`    // 分享员类型
	ShareScene    int    `json:"share_scene"`    // 分享场景
	SKUID         int64  `json:"sku_id"`         // 商品skuid
	FromWecom     bool   `json:"from_wecom"`     // 是否来自企微分享
}

// AgentInfo 授权账号信息
type AgentInfo struct {
	AgentFinderID       string `json:"agent_finder_id"`       // 授权视频号id
	AgentFinderNickname string `json:"agent_finder_nickname"` // 授权视频号昵称
}

// SourceInfo 订单来源信息
type SourceInfo struct {
	SKUID                  string `json:"sku_id"`                    // 商品skuid
	AccountType            int    `json:"account_type"`              // 带货账户类型
	AccountID              string `json:"account_id"`                // 带货账户id
	SaleChannel            int    `json:"sale_channel"`              // 销售渠道
	AccountNickname        string `json:"account_nickname"`          // 带货账户昵称
	ContentType            int    `json:"content_type"`              // 带货内容类型
	ContentID              string `json:"content_id"`                // 带货内容id
	PromoterHeadSupplierID string `json:"promoter_head_supplier_id"` // 自营推客推广的带货机构id
	OriginalID             string `json:"original_id"`               // 公众号/服务号 id
}

// RefundInfo 订单退款信息
type RefundInfo struct {
	RefundFreight int64 `json:"refund_freight"` // 退还运费金额
}

// GreetingCardInfo 需代写的商品贺卡信息
type GreetingCardInfo struct {
	GiverName       string `json:"giver_name"`       // 贺卡落款
	ReceiverName    string `json:"receiver_name"`    // 贺卡称谓
	GreetingMessage string `json:"greeting_message"` // 贺卡内容
}

// CustomInfo 商品定制信息
type CustomInfo struct {
	CustomImgURL        string `json:"custom_img_url"`         // 定制图片
	CustomWord          string `json:"custom_word"`            // 定制文字
	CustomType          int    `json:"custom_type"`            // 定制类型
	CustomPreviewImgURL string `json:"custom_preview_img_url"` // 定制预览图片
}

// AftersaleDetail 售后信息
type AftersaleDetail struct {
	AftersaleOrderList  []AftersaleOrderInfo `json:"aftersale_order_list"`   // 售后单列表
	OnAftersaleOrderCnt int                  `json:"on_aftersale_order_cnt"` // 正在售后流程的售后单数
}

// AftersaleOrderInfo 售后单信息
type AftersaleOrderInfo struct {
	AftersaleOrderID string `json:"aftersale_order_id"` // 售后单ID
	Status           int    `json:"status"`             // 售后单状态
}

// OrderPresentInfo 订单对应礼物单信息
type OrderPresentInfo struct {
	PresentNote         string `json:"present_note"`          // 礼物订单留言
	PresentGiverOpenID  string `json:"present_giver_openid"`  // 礼物订单赠送者openid
	PresentGiverUnionID string `json:"present_giver_unionid"` // 礼物订单赠送者在开放平台的唯一标识符
	PresentOrderIDStr   string `json:"present_order_id_str"`  // 礼物订单ID
	PresentSendType     int    `json:"present_send_type"`     // 礼物单类型
	IsB2CFreePresent    bool   `json:"is_b2c_free_present"`   // 礼物单是否付款
}

// IntraCityOrderInfo 同城单信息
type IntraCityOrderInfo struct {
	ShopID                 string `json:"shop_id"`                   // 门店id
	PredictArriveStartTime int64  `json:"predict_arrive_start_time"` // 预计送达开始时间
	PredictArriveEndTime   int64  `json:"predict_arrive_end_time"`   // 预计送达结束时间
	PredictArriveTimeType  int    `json:"predict_arrive_time_type"`  // 配送类型
}
