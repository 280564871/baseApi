package models

type OrderExpress struct {
	Id           int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	OrderId      int    `json:"order_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	ShipperId    int    `json:"shipper_id" xorm:"not null default 0 MEDIUMINT(8)"`
	ShipperName  string `json:"shipper_name" xorm:"not null default '' comment('物流公司名称') VARCHAR(120)"`
	ShipperCode  string `json:"shipper_code" xorm:"not null default '' comment('物流公司代码') VARCHAR(60)"`
	LogisticCode string `json:"logistic_code" xorm:"not null default '' comment('快递单号') VARCHAR(20)"`
	Traces       string `json:"traces" xorm:"not null default '' comment('物流跟踪信息') VARCHAR(2000)"`
	IsFinish     int    `json:"is_finish" xorm:"not null default 0 TINYINT(1)"`
	RequestCount int    `json:"request_count" xorm:"default 0 comment('总查询次数') INT(11)"`
	RequestTime  int    `json:"request_time" xorm:"default 0 comment('最近一次向第三方查询物流信息时间') INT(11)"`
	AddTime      int    `json:"add_time" xorm:"not null default 0 comment('添加时间') INT(11)"`
	UpdateTime   int    `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(11)"`
}
