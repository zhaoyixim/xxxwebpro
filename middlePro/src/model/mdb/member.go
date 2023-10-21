package mdb

import (
	"github.com/golang-module/carbon/v2"
)

type MemberModel struct {
	ID            int         `gorm:"column:id;primary_key" json:"id"`
	Cuname        string      `gorm:"column:cuname" json:"cuname"`
	Nickname      string      `gorm:"column:nickname" json:"nickname"`
	Passwd        string      `gorm:"column:passwd" json:"passwd"`
	Avatar        string      `gorm:"column:avatar" json:"avatar"`
	Ip            string      `gorm:"column:ip" json:"ip"`
	Sta           int         `gorm:"column:sta" json:"sta"`
	TotalU        string      `gorm:"column:total_U" json:"total_U"`
	Email         string      `gorm:"column:email" json:"email"`
	Msgcode       string      `gorm:"column:msgcode" json:"msgcode"`
	WalletAddress string      `gorm:"column:wallet_address" json:"wallet_address"`
	LoginTimes    int         `gorm:"column:login_times" json:"login_times"`
	UpdateAt      carbon.Time `gorm:"column:update_at" json:"update_at"`
	IsMech        int         `gorm:"column:is_mech" json:"is_mech"`
	TurnsTime     int         `gorm:"column:turns_time" json:"turns_time"`
}

func (m *MemberModel) TableName() string {
	return "member"
}
