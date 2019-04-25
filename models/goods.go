package models

type Goods struct {
	Id                int    `json:"id" xorm:"not null pk INT(11)"`
	CategoryId        int    `json:"category_id" xorm:"not null default 0 index INT(11)"`
	GoodsSn           string `json:"goods_sn" xorm:"not null default '' index VARCHAR(60)"`
	Name              string `json:"name" xorm:"not null default '' VARCHAR(120)"`
	BrandId           int    `json:"brand_id" xorm:"not null default 0 index INT(11)"`
	GoodsNumber       int    `json:"goods_number" xorm:"not null default 0 index MEDIUMINT(8)"`
	Keywords          string `json:"keywords" xorm:"not null default '' VARCHAR(255)"`
	GoodsBrief        string `json:"goods_brief" xorm:"not null default '' VARCHAR(255)"`
	GoodsDesc         string `json:"goods_desc" xorm:"TEXT"`
	IsOnSale          int    `json:"is_on_sale" xorm:"not null default 1 TINYINT(1)"`
	AddTime           int    `json:"add_time" xorm:"not null default 0 INT(10)"`
	SortOrder         int    `json:"sort_order" xorm:"not null default 100 index SMALLINT(4)"`
	IsDelete          int    `json:"is_delete" xorm:"not null default 0 TINYINT(1)"`
	AttributeCategory int    `json:"attribute_category" xorm:"not null default 0 INT(11)"`
	CounterPrice      string `json:"counter_price" xorm:"not null default 0.00 comment('专柜价格') DECIMAL(10,2)"`
	ExtraPrice        string `json:"extra_price" xorm:"not null default 0.00 comment('附加价格') DECIMAL(10,2)"`
	IsNew             int    `json:"is_new" xorm:"not null default 0 TINYINT(1)"`
	GoodsUnit         string `json:"goods_unit" xorm:"not null comment('商品单位') VARCHAR(45)"`
	PrimaryPicUrl     string `json:"primary_pic_url" xorm:"not null comment('商品主图') VARCHAR(255)"`
	ListPicUrl        string `json:"list_pic_url" xorm:"not null comment('商品列表图') VARCHAR(255)"`
	RetailPrice       string `json:"retail_price" xorm:"not null default 0.00 comment('零售价格') DECIMAL(10,2)"`
	SellVolume        int    `json:"sell_volume" xorm:"not null default 0 comment('销售量') INT(11)"`
	PrimaryProductId  int    `json:"primary_product_id" xorm:"not null default 0 comment('主sku　product_id') INT(11)"`
	UnitPrice         string `json:"unit_price" xorm:"not null default 0.00 comment('单位价格，单价') DECIMAL(10,2)"`
	PromotionDesc     string `json:"promotion_desc" xorm:"not null VARCHAR(255)"`
	PromotionTag      string `json:"promotion_tag" xorm:"not null VARCHAR(45)"`
	AppExclusivePrice string `json:"app_exclusive_price" xorm:"not null comment('APP专享价') DECIMAL(10,2)"`
	IsAppExclusive    int    `json:"is_app_exclusive" xorm:"not null comment('是否是APP专属') TINYINT(1)"`
	IsLimited         int    `json:"is_limited" xorm:"not null TINYINT(1)"`
	IsHot             int    `json:"is_hot" xorm:"not null default 0 TINYINT(1)"`
}
