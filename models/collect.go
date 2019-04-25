package models

type Collect struct {
	Id          int `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	UserId      int `json:"user_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	ValueId     int `json:"value_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	AddTime     int `json:"add_time" xorm:"not null default 0 INT(11)"`
	IsAttention int `json:"is_attention" xorm:"not null default 0 comment('是否是关注') index TINYINT(1)"`
	TypeId      int `json:"type_id" xorm:"not null default 0 INT(2)"`
}
