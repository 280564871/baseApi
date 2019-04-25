package models

type Shipper struct {
	Id        int    `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	Name      string `json:"name" xorm:"not null default '' comment('快递公司名称') VARCHAR(20)"`
	Code      string `json:"code" xorm:"not null default '' comment('快递公司代码') VARCHAR(10)"`
	SortOrder int    `json:"sort_order" xorm:"not null default 10 comment('排序') INT(11)"`
}
