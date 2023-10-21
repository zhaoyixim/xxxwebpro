package mdb

type BotSettingModel struct {
	ID        int    `gorm:"column:id;primary_key" json:"id"`
	Tgtoken   string `gorm:"column:tgtoken" json:"tgtoken"`
	Belongs   int    `gorm:"column:belongs" json:"belongs"`
	Cname     string `gorm:"column:cname" json:"cname"`
	Cusername string `gorm:"column:cusername" json:"cusername"`
	MechId    int    `gorm:"column:mech_id" json:"mech_id"`
	Ctype     int    `gorm:"column:ctype" json:"ctype"`
	Sta       int    `gorm:"column:sta" json:"sta"`
}

// TableName sets the insert table name for this struct type
func (c *BotSettingModel) TableName() string {
	return "bot_setting"
}
