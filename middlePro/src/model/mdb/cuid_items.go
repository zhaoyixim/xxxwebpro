package mdb

type CuidItemModel struct {
	ID       int `gorm:"column:id;primary_key" json:"id"`
	Cuid     int `gorm:"column:cuid" json:"cuid"`
	Sta      int `gorm:"column:sta" json:"sta"`
	ItemId   int `gorm:"column:item_id" json:"item_id"`
	CardN    int `gorm:"column:card_N" json:"card_N"`
	RemainN  int `gorm:"column:remain_N" json:"remain_N"`
	IsLimitN int `gorm:"column:is_limit_N" json:"is_limit_N"`
}

// TableName sets the insert table name for this struct type
func (c *CuidItemModel) TableName() string {
	return "cuid_items"
}
