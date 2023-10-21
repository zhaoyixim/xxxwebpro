package mdb

import (
	"github.com/golang-module/carbon/v2"
)

type GameItemList struct {
	Cuid        int         `gorm:"column:cuid" json:"cuid"`
	ItemId      int         `gorm:"column:item_id" json:"item_id"`
	PayUnit     int         `gorm:"column:pay_unit" json:"pay_unit"`
	TgKefu      string      `gorm:"column:tg_kefu" json:"tg_kefu"`
	TgGroup     string      `gorm:"column:tg_group" json:"tg_group"`
	MechId      int         `gorm:"column:mech_id" json:"mech_id"`
	Cuname      string      `gorm:"column:cuname" json:"cuname"`
	Cavatar     string      `gorm:"column:cavatar" json:"cavatar"`
	Ctitle      string      `gorm:"column:ctitle" json:"ctitle"`
	TotalU      int         `gorm:"column:total_U" json:"total_U"`
	ChargedU    int         `gorm:"column:charged_U" json:"charged_U"`
	PayCount    int         `gorm:"column:pay_count" json:"pay_count"`
	IsReword    int         `gorm:"column:is_reword" json:"is_reword"` // 0：没中奖-默认 1:中奖
	Breif       string      `gorm:"column:breif" json:"breif"`
	RewordAll   string      `gorm:"column:reword_all" json:"reword_all"` // 所有中奖信息
	Sta         int         `gorm:"column:sta" json:"sta"`
	ID          int         `gorm:"column:id;primary_key" json:"id"`
	RewordCount int         `gorm:"column:reword_count" json:"reword_count"`
	CreatedAt   carbon.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   carbon.Time `gorm:"column:updated_at" json:"updated_at"`
	Reword1     int         `gorm:"column:reword1" json:"reword1"`
	Reword2     int         `gorm:"column:reword2" json:"reword2"`
	Reword3     int         `gorm:"column:reword3" json:"reword3"`
	Reword4     int         `gorm:"column:reword4" json:"reword4"`
	Reword5     int         `gorm:"column:reword5" json:"reword5"`
	RewordNum1  int         `gorm:"column:reword_num1" json:"reword_num1"`
	RewordNum2  int         `gorm:"column:reword_num2" json:"reword_num2"`
	RewordNum3  int         `gorm:"column:reword_num3" json:"reword_num3"`
	RewordNum4  int         `gorm:"column:reword_num4" json:"reword_num4"`
	RewordNum5  int         `gorm:"column:reword_num5" json:"reword_num5"`
}

func (g *GameItemList) TableName() string {
	return "game_item_list"
}
