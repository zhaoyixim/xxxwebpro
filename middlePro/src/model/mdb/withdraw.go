package mdb

import (
	"github.com/golang-module/carbon/v2"
)

type WithdrawModel struct {
	ID            int         `gorm:"column:id;primary_key" json:"id"`
	Cuid          int         `gorm:"column:cuid" json:"cuid"`
	WalletAddress string      `gorm:"column:wallet_address" json:"wallet_address"`
	CreatedAt     carbon.Time `gorm:"column:create_at" json:"create_at"`
	FromAddress   string      `gorm:"column:from_address" json:"from_address"`
	ToAddress     string      `gorm:"column:to_address" json:"to_address"`
	Ctype         int         `gorm:"column:ctype" json:"ctype"`
	Amount        int         `gorm:"column:amount" json:"amount"`
	Rate          int         `gorm:"column:rate" json:"rate"`
	WithdrawU     string      `gorm:"column:withdraw_U" json:"withdraw_U"`
	Sta           int         `gorm:"column:sta" json:"sta"`
}

func (m *WithdrawModel) TableName() string {
	return "withdraw"
}
