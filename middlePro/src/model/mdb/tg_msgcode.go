package mdb

import "github.com/golang-module/carbon/v2"

type MsgCode struct {
	ID          int         `gorm:"column:id;primary_key" json:"id"`
	Msgcode     string      `gorm:"column:msg_code" json:"msg_code"`
	TgId        int64       `gorm:"column:tg_id" json:"tg_id"`
	TgFirstname string      `gorm:"column:tg_firstname" json:"tg_firstname"`
	TgLastname  string      `gorm:"column:tg_lastname" json:"tg_lastname"`
	TgUsername  string      `gorm:"column:tg_username" json:"tg_username"`
	TgIsbot     int         `gorm:"column:tg_isbot" json:"tg_isbot"`
	CreateAt    carbon.Time `gorm:"column:create_at" json:"create_at"`
	Sta         int         `gorm:"column:sta" json:"sta"`
}

// TableName sets the insert table name for this struct type
func (m *MsgCode) TableName() string {
	return "tg_msgcode"
}
