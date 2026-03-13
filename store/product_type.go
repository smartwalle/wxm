package store

import "github.com/smartwalle/wxm"

// ProductStatus 商品状态
type ProductStatus int

const (
	ProductStatusInitial        ProductStatus = 0  // 初始值
	ProductStatusEditing        ProductStatus = 1  // 编辑中
	ProductStatusAuditing       ProductStatus = 2  // 审核中
	ProductStatusAuditFailed    ProductStatus = 3  // 审核失败
	ProductStatusAuditSuccess   ProductStatus = 4  // 审核成功
	ProductStatusOnShelf        ProductStatus = 5  // 上架
	ProductStatusRecycle        ProductStatus = 6  // 回收站
	ProductStatusUploading      ProductStatus = 7  // 商品异步提交，上传中
	ProductStatusUploadFailed   ProductStatus = 8  // 商品异步提交，上传失败
	ProductStatusDeleted        ProductStatus = 9  // 彻底删除，商品无法再进行任何操作
	ProductStatusFrozen         ProductStatus = 10 // 冻结, 审核通过但是不能上架
	ProductStatusOffShelf       ProductStatus = 11 // 自主下架
	ProductStatusSoldOut        ProductStatus = 12 // 售罄下架
	ProductStatusViolationOff   ProductStatus = 13 // 违规下架/风控系统下架
	ProductStatusDepositOff     ProductStatus = 14 // 保证金不足下架
	ProductStatusBrandExpired   ProductStatus = 15 // 品牌过期下架
	ProductStatusBanned         ProductStatus = 20 // 商品被封禁
	ProductStatusSKULogicDelete ProductStatus = 21 // sku逻辑删除
	ProductStatusNotExist       ProductStatus = 30 // 商品不存在
	ProductQualityCheckFailed   ProductStatus = 71 // 质检不通过
)

// EditStatus 商品草稿状态
type EditStatus int

const (
	EditStatusInitial      EditStatus = 0 // 初始值
	EditStatusEditing      EditStatus = 1 // 编辑中
	EditStatusAuditing     EditStatus = 2 // 审核中
	EditStatusAuditFailed  EditStatus = 3 // 审核失败
	EditStatusAuditSuccess EditStatus = 4 // 审核成功
	EditStatusUploading    EditStatus = 7 // 商品异步提交，上传中
	EditStatusUploadFailed EditStatus = 8 // 商品异步提交，上传失败
)

// DataType 数据类型
type DataType int

const (
	DataTypeOnline      DataType = 1 // 获取线上数据
	DataTypeDraft       DataType = 2 // 获取草稿数据
	DataTypeOnlineDraft DataType = 3 // 同时获取线上和草稿数据（注意：上架过的商品才有线上数据）
)

// GetProductListRequest 获取商品列表请求参数
type GetProductListRequest struct {
	Status   ProductStatus `json:"status,omitempty"`   // 商品状态，不填默认拉全部商品（不包含从未上架的草稿和回收站商品）
	PageSize int           `json:"page_size"`          // 每页数量（默认10，不超过30）
	NextKey  string        `json:"next_key,omitempty"` // 由上次请求返回，记录翻页的上下文
}

// GetProductListResponse 获取商品列表响应
type GetProductListResponse struct {
	wxm.Error
	ProductIDs []string `json:"product_ids"` // 商品id列表
	NextKey    string   `json:"next_key"`    // 本次翻页的上下文，用于请求下一页
	TotalNum   int      `json:"total_num"`   // 商品总数
}

// GetProductRequest 获取商品请求参数
type GetProductRequest struct {
	ProductID string   `json:"product_id"`          // 商品ID
	DataType  DataType `json:"data_type,omitempty"` // 数据类型：1-获取线上数据；2-获取草稿数据；3-同时获取线上和草稿数据
}

// GetProductResponse 获取商品响应
type GetProductResponse struct {
	wxm.Error
	Product       *Product       `json:"product,omitempty"`         // 商品线上数据
	EditProduct   *Product       `json:"edit_product,omitempty"`    // 商品草稿数据
	SaleLimitInfo *SaleLimitInfo `json:"sale_limit_info,omitempty"` // 当日售卖上限提醒
	InfoScore     *InfoScore     `json:"info_score,omitempty"`      // 商品信息质量
	CmpPriceInfo  *CmpPriceInfo  `json:"cmp_price_info,omitempty"`  // 商品高价预警
	AuditInfo     *AuditInfo     `json:"audit_info,omitempty"`      // 审核信息
}

// Product 商品信息
type Product struct {
	ProductID        string            `json:"product_id"`                   // 小店内部商品ID
	OutProductID     string            `json:"out_product_id,omitempty"`     // 外部平台自定义商品ID
	Title            string            `json:"title,omitempty"`              // 标题
	SubTitle         string            `json:"sub_title,omitempty"`          // 副标题（已废弃）
	HeadImgs         []string          `json:"head_imgs,omitempty"`          // 主图，多张，列表，最多9张，每张不超过2MB
	DescInfo         *DescInfo         `json:"desc_info,omitempty"`          // 商品详情
	DeliverMethod    int               `json:"deliver_method,omitempty"`     // 发货方式：0-快递发货；1-无需快递，手机号发货；3-无需快递，可选发货账号类型
	DeliverAcctType  []int             `json:"deliver_acct_type,omitempty"`  // 发货账号：1-微信openid；2-QQ号；3-手机号；4-邮箱
	ExpressInfo      *ExpressInfo      `json:"express_info,omitempty"`       // 运费信息
	AftersaleDesc    string            `json:"aftersale_desc,omitempty"`     // 售后说明
	LimitedInfo      *LimitedInfo      `json:"limited_info,omitempty"`       // 限购信息
	ExtraService     *ExtraService     `json:"extra_service,omitempty"`      // 额外服务
	Status           ProductStatus     `json:"status,omitempty"`             // 商品线上状态
	EditStatus       EditStatus        `json:"edit_status,omitempty"`        // 商品草稿状态
	MinPrice         int64             `json:"min_price,omitempty"`          // 商品 SKU 最小价格（单位：分）
	Cats             []Cat             `json:"cats,omitempty"`               // 商品类目（旧）
	CatsV2           []Cat             `json:"cats_v2,omitempty"`            // 新类目树
	Attrs            []Attribute       `json:"attrs,omitempty"`              // 属性键key（属性自定义用）
	SPUCode          string            `json:"spu_code,omitempty"`           // 商家自定义的商品编码
	BrandID          string            `json:"brand_id,omitempty"`           // 品牌id，无品牌为"2100000000"
	SKUs             []SKU             `json:"skus,omitempty"`               // sku信息
	ProductType      int               `json:"product_type,omitempty"`       // 商品类型。1: 小店普通自营商品；2: 福袋抽奖商品；3: 直播间闪电购商品
	EditTime         int64             `json:"edit_time,omitempty"`          // 商品草稿最近一次修改时间
	AfterSaleInfo    *AfterSaleInfo    `json:"after_sale_info,omitempty"`    // 商品售后信息
	SrcProductID     string            `json:"src_product_id,omitempty"`     // 来源商品id
	ProductQuaInfos  []ProductQuaInfo  `json:"product_qua_infos,omitempty"`  // 商品资质列表
	SizeChart        *SizeChart        `json:"size_chart,omitempty"`         // 尺码信息
	TimingOnsaleInfo *TimingOnsaleInfo `json:"timing_onsale_info,omitempty"` // 商品待开售信息
	ShortTitle       string            `json:"short_title,omitempty"`        // 短标题
	TotalSoldNum     int64             `json:"total_sold_num,omitempty"`     // 销量
	ReleaseMode      int               `json:"release_mode,omitempty"`       // 发布模式，0: 普通模式；1: 极简模式
	SPUDeliverInfo   *SPUDeliverInfo   `json:"spu_deliver_info,omitempty"`   // spu维度的sku预售配置
}

// Cat 类目
type Cat struct {
	CatID string `json:"cat_id"` // 类目ID
}

// Attribute 属性
type Attribute struct {
	AttrKey   string `json:"attr_key,omitempty"`   // 属性键key
	AttrValue string `json:"attr_value,omitempty"` // 属性值
}

// DescInfo 商品详情
type DescInfo struct {
	Imgs []string `json:"imgs,omitempty"` // 商品详情图片(最多20张)
	Desc string   `json:"desc,omitempty"` // 商品详情文字
}

// ExpressInfo 运费信息
type ExpressInfo struct {
	TemplateID string `json:"template_id,omitempty"` // 运费模板ID
	Weight     int64  `json:"weight,omitempty"`      // 商品重量，单位克
}

// LimitedInfo 限购信息
type LimitedInfo struct {
	PeriodType    int `json:"period_type,omitempty"`     // 限购周期类型，0无限购，1按自然日限购，2按自然周限购，3按自然月限购
	LimitedBuyNum int `json:"limited_buy_num,omitempty"` // 限购数量
}

// ExtraService 额外服务
type ExtraService struct {
	SevenDayReturn   int `json:"seven_day_return,omitempty"`   // 是否支持七天无理由退货
	PayAfterUse      int `json:"pay_after_use,omitempty"`      // 是否支持先用后付
	FreightInsurance int `json:"freight_insurance,omitempty"`  // 是否支持运费险
	DamageGuarantee  int `json:"damage_guarantee,omitempty"`   // 是否支持假一赔三
	FakeOnePayThree  int `json:"fake_one_pay_three,omitempty"` // 是否支持坏损包退
	ExchangeSupport  int `json:"exchange_support,omitempty"`   // 是否支持换货
}

// SKU 商品SKU信息
type SKU struct {
	SKUID          string          `json:"sku_id"`                     // skuID
	OutSKUID       string          `json:"out_sku_id,omitempty"`       // 外部平台自定义skuID
	ThumbImg       string          `json:"thumb_img,omitempty"`        // sku小图
	SalePrice      int64           `json:"sale_price,omitempty"`       // 售卖价格，以分为单位
	StockNum       int             `json:"stock_num,omitempty"`        // sku库存
	SKUCode        string          `json:"sku_code,omitempty"`         // 商家自定义的sku编码
	SKUAttrs       []Attribute     `json:"sku_attrs,omitempty"`        // sku属性
	Status         ProductStatus   `json:"status,omitempty"`           // sku状态
	SKUDeliverInfo *SKUDeliverInfo `json:"sku_deliver_info,omitempty"` // sku库存情况
	BarCode        string          `json:"bar_code,omitempty"`         // sku条形码
}

// SKUDeliverInfo SKU库存情况
type SKUDeliverInfo struct {
	StockType                      int   `json:"stock_type,omitempty"`                         // sku库存情况。0:现货（默认），1:全款预售
	FullPaymentPresaleDeliveryType int   `json:"full_payment_presale_delivery_type,omitempty"` // sku发货节点，该字段仅对stock_type=1有效
	PresaleBeginTime               int64 `json:"presale_begin_time,omitempty"`                 // sku预售周期开始时间，秒级时间戳
	PresaleEndTime                 int64 `json:"presale_end_time,omitempty"`                   // sku预售周期结束时间，秒级时间戳
	FullPaymentPresaleDeliveryTime int   `json:"full_payment_presale_delivery_time,omitempty"` // sku发货时效
	SpotAfterPresaleEnd            int   `json:"spot_after_presale_end,omitempty"`             // 是否在预售结束后自动转为现货
	PredictDeliveryTime            int64 `json:"predict_delivery_time"`                        // 预计发货时间(stock_type=1时返回该字段)
}

// AfterSaleInfo 商品售后信息
type AfterSaleInfo struct {
	AfterSaleAddressID string `json:"after_sale_address_id,omitempty"` // 商品的售后地址id
}

// ProductQuaInfo 商品资质
type ProductQuaInfo struct {
	QuaID  int64    `json:"qua_id"`            // 商品资质id
	QuaURL []string `json:"qua_url,omitempty"` // 商品资质图片列表
}

// SizeChart 尺码信息
type SizeChart struct {
	Enable            bool            `json:"enable"`                       // 是否启用尺码表
	SpecificationList []Specification `json:"specification_list,omitempty"` // 尺码表
}

// Specification 尺码表规格
type Specification struct {
	Name      string      `json:"name,omitempty"`       // 尺码属性名称
	Unit      string      `json:"unit,omitempty"`       // 尺码属性值的单位
	IsRange   bool        `json:"is_range,omitempty"`   // 尺码属性值是否为区间
	ValueList []ValueInfo `json:"value_list,omitempty"` // 尺码值与尺码属性值的映射列表
}

// ValueInfo 尺码值信息
type ValueInfo struct {
	Key   string `json:"key,omitempty"`   // 尺码值
	Value string `json:"value,omitempty"` // 尺码属性值
	Left  string `json:"left,omitempty"`  // 尺码属性值的左边界
	Right string `json:"right,omitempty"` // 尺码属性值的右边界
}

// TimingOnsaleInfo 商品待开售信息
type TimingOnsaleInfo struct {
	Status      int   `json:"status,omitempty"`        // 状态
	OnsaleTime  int64 `json:"onsale_time,omitempty"`   // 开售时间，秒级时间戳，0为未配置时间
	IsHidePrice int   `json:"is_hide_price,omitempty"` // 是否隐藏价格
	TaskID      int64 `json:"task_id,omitempty"`       // 待开售任务ID
}

// SPUDeliverInfo spu维度的sku预售配置
type SPUDeliverInfo struct {
	SKUDeliverInfo *SKUDeliverInfo `json:"sku_deliver_info,omitempty"` // sku预售配置
	IsSPURange     int             `json:"is_spu_range,omitempty"`     // 是否生效
}

// SaleLimitInfo 当日售卖上限提醒
type SaleLimitInfo struct {
	IsLimited int    `json:"is_limit,omitempty"`  // 是否受到管控
	Title     string `json:"title,omitempty"`     // 售卖限制标题
	SubTitle  string `json:"sub_title,omitempty"` // 售卖限制描述
}

// InfoScore 商品信息质量
type InfoScore struct {
	ScoreLevel   int        `json:"score_level,omitempty"`    // 商品信息质量总分等级
	SubScoreList []SubScore `json:"sub_score_list,omitempty"` // 商品信息质量子项
}

// SubScore 商品信息质量子项
type SubScore struct {
	AuditRemark      string `json:"audit_remark,omitempty"`        // 审核备注
	BusiDataFieldAPI string `json:"busi_data_field_api,omitempty"` // 商品信息字段名
	FieldName        string `json:"field_name,omitempty"`          // 字段描述
}

// CmpPriceInfo 商品高价预警
type CmpPriceInfo struct {
	Status        int         `json:"status,omitempty"`          // 检测状态
	SKUResultList []SKUResult `json:"sku_result_list,omitempty"` // sku高价预警检测结果
}

// SKUResult SKU高价预警结果
type SKUResult struct {
	SKUID              int64 `json:"sku_id"`               // skuID
	IsExpensiveWarning bool  `json:"is_expensive_warning"` // 是否被检测为高价
}

// AuditInfo 审核信息
type AuditInfo struct {
	UserStrategyFlagList []int `json:"user_strategy_flag_list,omitempty"` // 提审的商品上架策略
}

// GetStockRequest 获取库存请求参数
type GetStockRequest struct {
	ProductID string `json:"product_id"` // 内部商品ID
	SKUID     string `json:"sku_id"`     // 内部sku_id
}

// GetStockResponse 获取库存响应
type GetStockResponse struct {
	wxm.Error
	Data *StockData `json:"data,omitempty"` // 库存数据
}

// StockData 库存数据
type StockData struct {
	NormalStockNum          int              `json:"normal_stock_num"`           // 通用库存数量
	LimitedDiscountStockNum int              `json:"limited_discount_stock_num"` // 限时抢购库存数量
	WarehouseStocks         []WarehouseStock `json:"warehouse_stocks,omitempty"` // 区域库存
	TotalStockNum           int              `json:"total_stock_num"`            // 库存总量：通用库存数量 + 限时抢购库存数量 + 区域库存总量
}

// WarehouseStock 区域库存
type WarehouseStock struct {
	OutWarehouseID string `json:"out_warehouse_id"` // 区域库存外部id
	Num            int    `json:"num"`              // 区域库存数量
	LockStock      int    `json:"lock_stock"`       // 区域库存的锁定库存（已下单未支付的库存）数量
}

// BatchGetStockRequest 批量获取库存信息请求参数
type BatchGetStockRequest struct {
	ProductIDs []string `json:"product_id"` // 商品ID列表，上限为50
}

// BatchGetStockResponse 批量获取库存信息响应
type BatchGetStockResponse struct {
	wxm.Error
	Data *BatchStockData `json:"data,omitempty"` // 批量库存数据
}

// BatchStockData 批量库存数据
type BatchStockData struct {
	SPUStockList []SPUStock `json:"spu_stock_list,omitempty"` // spu库存
}

// SPUStock spu库存
type SPUStock struct {
	ProductID string     `json:"product_id"`          // 商品ID
	SKUStock  []SKUStock `json:"sku_stock,omitempty"` // sku库存
}

// SKUStock sku库存
type SKUStock struct {
	SKUID                   string           `json:"sku_id"`                     // skuID
	NormalStockNum          int              `json:"normal_stock_num"`           // 普通/通用库存数量
	LimitedDiscountStockNum int              `json:"limited_discount_stock_num"` // 限时抢购库存数量
	WarehouseStocks         []WarehouseStock `json:"warehouse_stocks,omitempty"` // 区域库存
	FinderTotalNum          int              `json:"finder_total_num,omitempty"` // 达人专属计划营销库存数量
	TotalStockNum           int              `json:"total_stock_num"`            // 库存总量：普通/通用库存数量 + 限时抢购库存数量 + 区域库存总量 + 直播预热/专享库存
	ExclusiveNum            int              `json:"exclusive_num,omitempty"`    // 直播预热专属库存
}
