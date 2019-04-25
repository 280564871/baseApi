package models

type Attribute struct {
	Id                  int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	AttributeCategoryId int    `json:"attribute_category_id" xorm:"not null default 0 index INT(11)"`
	Name                string `json:"name" xorm:"not null default '' VARCHAR(60)"`
	InputType           int    `json:"input_type" xorm:"not null default 1 TINYINT(1)"`
	Values              string `json:"values" xorm:"not null TEXT"`
	SortOrder           int    `json:"sort_order" xorm:"not null default 0 TINYINT(3)"`
}
