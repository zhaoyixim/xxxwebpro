package response

import "github.com/pengk/summer/model/mdb"

type CheckoutCounterResponse struct {
	TradeId        string  `json:"trade_id"`        //  pengk订单号
	ActualAmount   float64 `json:"actual_amount"`   //  订单实际需要支付的金额，保留4位小数
	Token          string  `json:"token"`           //  收款钱包地址
	ExpirationTime int64   `json:"expiration_time"` // 过期时间 时间戳
	RedirectUrl    string  `json:"redirect_url"`
}

type CheckStatusResponse struct {
	TradeId string `json:"trade_id"` //  pengk订单号
	Status  int    `json:"status"`
}

type ItemTimesReturns struct {
	Message     string            `json:"message"` //  pengk订单号
	Status      int               `json:"status"`  // 1--成功 2 --失败返回s
	RestTimes   int               `json:"rest_times"`
	OrdersData  []mdb.Orders      `json:"item_data"`
	MembersData []mdb.MemberModel `json:"member_data"`
	ItemSta     int               `json:"item_sta"`
}
