package mdb

type SysConfig struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	Cname  string `gorm:"column:cname" json:"cname"`
	Ckey   string `gorm:"column:ckey" json:"ckey"`
	Cvalue string `gorm:"column:cvalue" json:"cvalue"`
	MechId int    `gorm:"column:mech_id" json:"mech_id"`
	Desp   string `gorm:"column:desp" json:"desp"`
}

func (n *SysConfig) TableName() string {
	return "sys_setting"
}
