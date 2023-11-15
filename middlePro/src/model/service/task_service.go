package service

import (
	"fmt"
	"sync"

	"github.com/hibiken/asynq"
	"github.com/pengk/summer/model/data"
	"github.com/pengk/summer/model/mdb"
	"github.com/pengk/summer/model/request"
	"github.com/pengk/summer/mq"
	"github.com/pengk/summer/mq/handle"
	"github.com/pengk/summer/util/json"
	"github.com/pengk/summer/util/log"
	"github.com/shopspring/decimal"
)

const UsdtTrc20ApiUri = "https://apilist.tronscanapi.com/api/transfer/trc20"

type UsdtTrc20Resp struct {
	PageSize int    `json:"page_size"`
	Code     int    `json:"code"`
	Data     []Data `json:"data"`
}

type TokenInfo struct {
	TokenID      string `json:"tokenId"`
	TokenAbbr    string `json:"tokenAbbr"`
	TokenName    string `json:"tokenName"`
	TokenDecimal int    `json:"tokenDecimal"`
	TokenCanShow int    `json:"tokenCanShow"`
	TokenType    string `json:"tokenType"`
	TokenLogo    string `json:"tokenLogo"`
	TokenLevel   string `json:"tokenLevel"`
	IssuerAddr   string `json:"issuerAddr"`
	Vip          bool   `json:"vip"`
}

type Data struct {
	Amount         string `json:"amount"`
	ApprovalAmount string `json:"approval_amount"`
	BlockTimestamp int64  `json:"block_timestamp"`
	Block          int    `json:"block"`
	From           string `json:"from"`
	To             string `json:"to"`
	Hash           string `json:"hash"`
	Confirmed      int    `json:"confirmed"`
	ContractType   string `json:"contract_type"`
	ContracTType   int    `json:"contractType"`
	Revert         int    `json:"revert"`
	ContractRet    string `json:"contract_ret"`
	EventType      string `json:"event_type"`
	IssueAddress   string `json:"issue_address"`
	Decimals       int    `json:"decimals"`
	TokenName      string `json:"token_name"`
	ID             string `json:"id"`
	Direction      int    `json:"direction"`
}

//回调

// Trc20CallBack trc20回调

// Trc20CallBack trc20回调
func NewTrc20CallBack(token string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			log.Sugar.Error(err)
		}
	}()

	// client := http_client.GetHttpClient()
	// startTime := carbon.Now().AddHours(-24).TimestampWithMillisecond()
	// endTime := carbon.Now().TimestampWithMillisecond()
	// resp, err := client.R().SetQueryParams(map[string]string{
	// 	"sort":            "-timestamp",
	// 	"reverse":         "true",
	// 	"limit":           "50",
	// 	"start":           "0",
	// 	"direction":       "2",
	// 	"db_version":      "1",
	// 	"trc20Id":         "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
	// 	"address":         token,
	// 	"start_timestamp": stdutil.ToString(startTime),
	// 	"end_timestamp":   stdutil.ToString(endTime),
	// }).Get(UsdtTrc20ApiUri)

	resp := `{"contractMap":{"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX":false,"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays":false,"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC":false},"tokenInfo":{"tokenId":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","tokenAbbr":"USDT","tokenName":"Tether USD","tokenDecimal":6,"tokenCanShow":1,"tokenType":"trc20","tokenLogo":"https://static.tronscan.org/production/logo/usdtlogo.png","tokenLevel":"2","issuerAddr":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","vip":true},"page_size":1,"code":200,"data":[{"amount":"100000000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2aaa07c7ff9e1cfe0fd2b69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100010000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2aaa07c7ff9e1cfe0fd2b69db8c3d14221f93452345","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100020000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2aaa07c7ff9e1cfe0fd2b69db8c3d1422fsgsert","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100030000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f74wertwer0fd2b69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100040000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411fwretwertwert342b69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100050000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb423wsrt9db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100060000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbesdfgsdfg45cb2aaa07c7ff9e1cfe0fd2b69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100070000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb2sdfty4c7ff9e1cfe0fd2b69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100080000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971sdftywe5r69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100090000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2sdfg5rshg69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"100100000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2asfdg3e4s9db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"1001100000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2aaasdfgsefrtb8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"1001200000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745sgsdrfts1cfe0fd2b69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"1001300000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2aasdfgs4tb69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"1001400000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f745cb2aasdf45t3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"1001500000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb237971411f74sg4t545tg2b69db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2},{"amount":"1001600000","status":0,"approval_amount":"0","block_timestamp":1695171189000,"block":54846815,"from":"TUeZ2fREiiGwLpAgZxd9H8NEsYqj8L8DKX","to":"TK1gJS6iEm5aL2U9jYKekHK4Rs4iEc5ays","hash":"cbeb23797141sdfgs56sdrysdr9db8c3d14221f963f526a8d","confirmed":1,"contract_type":"TriggerSmartContract","contractType":31,"revert":0,"contract_ret":"SUCCESS","event_type":"Transfer","issue_address":"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC","decimals":6,"token_name":"TetherToken","id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","direction":2}]}`
	// if err != nil {
	// 	panic(err)
	// }
	// if resp.StatusCode() != http.StatusOK {
	// 	panic(err)
	// }
	var trc20Resp UsdtTrc20Resp
	//err = json.Cjson.Unmarshal(resp.Body(), &trc20Resp)
	err := json.Cjson.Unmarshal([]byte(resp), &trc20Resp)
	if err != nil {
		panic(err)
	}
	// if trc20Resp.PageSize <= 0 {
	// 	return
	// }

	for _, transfer := range trc20Resp.Data {
		// if transfer.To != token || transfer.ContractRet != "SUCCESS" {
		// 	continue
		// }
		decimalQuant, err := decimal.NewFromString(transfer.Amount)
		if err != nil {
			panic(err)
		}
		decimalDivisor := decimal.NewFromFloat(1000000)
		amount := decimalQuant.Div(decimalDivisor).InexactFloat64()
		tradeId, err := data.GetTradeIdByWalletAddressAndAmount(token, amount)
		if err != nil {
			panic(err)
		}
		if tradeId == "" {
			continue
		}
		order, err := data.GetOrderInfoByTradeId(tradeId)
		fmt.Println("======33333333==========")
		fmt.Println(tradeId)
		fmt.Println(order.Status)
		if order.Status != 1 {
			continue
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("======444444==========")
		// 区块的确认时间必须在订单创建时间之后
		//createTime := order.CreatedAt.TimestampWithMillisecond()
		// if transfer.BlockTimestamp < createTime {
		// 	panic("Orders cannot actually be matched")
		// }
		// 到这一步就完全算是支付成功了
		req := &request.OrderProcessingRequest{
			Token:              token,
			TradeId:            tradeId,
			Amount:             amount,
			ItemId:             order.ItemId,
			BlockTransactionId: transfer.Hash,
			SCuid:              order.Cuid,
			OrderId:            int(order.ID),
		}
		//处理订单的收支问题
		//订单信息- order
		// relizeAmount := math.Floor(amount)
		// targetOrderLen := len(GetOrders)
		// findItemCuidData, _ := data.GetLimitNum(order.ItemId, order.Cuid)

		//默认模式-第一模式

		// if err != nil {
		// 	panic(err)
		// }

		//GItemInfo, _ := data.GetGDtlById(order.ItemId)

		GItemList := OrderProcessing(req, order)

		GetOrders, _ := data.GetOrderDatasByItemId(order.ItemId)
		gameOrderNum := len(GetOrders)
		fmt.Println("=====gameOrderNum====")
		fmt.Println(gameOrderNum)

		// if gameOrderNum == 0 {
		// 	gameOrderNum = 1
		// }
		//resNum := GItemInfo.RewordCount - targetOrderLength

		// if err != nil {
		// 	panic(err)
		// }
		//插入withdraw 进账
		InsertWithdraw := &mdb.WithdrawModel{
			Cuid:        order.Cuid,
			Amount:      int(amount),
			Ctype:       1,
			FromAddress: transfer.From,
			ToAddress:   transfer.To,
		}
		data.SetInAccountResult(InsertWithdraw)
		// 回调队列
		orderCallbackQueue, _ := handle.NewOrderCallbackQueue(order)
		mq.MClient.Enqueue(orderCallbackQueue, asynq.MaxRetry(5))
		msgTpl := ""

		// 发送机器人消息
		msgTpl = `
<b>新单号支付成功！</b>
<pre>交易号：%s</pre>
<pre>实际支付金额：%f usdt</pre>
<pre>钱包地址：%s</pre>
<pre>您的游戏顺序号为:%d</pre>
<pre>等待开奖</pre>
`
		msg := fmt.Sprintf(msgTpl, order.TradeId, order.ActualAmount, order.Token, gameOrderNum)

		fmt.Println(msg)
		//telegram.SendToBot(msg)
		findLimitN, _ := data.GetLimitNumByGListId(order.ItemId, order.Cuid)
		if GItemList.Sta == 2 {
			if findLimitN.IsLimitN == 0 {
				//第一种
				msgTpl = `				
					<pre>开奖条件满足，准备开奖</pre>
					<pre>中奖信息如下:%s</pre>				
					`
				msg := fmt.Sprintf(msgTpl, GItemList.RewordAll)
				fmt.Println(msg)
				//telegram.SendToBot(msg)

			}
			//第二种不用管-手动停止
			if findLimitN.IsLimitN == 1 {
				msgTpl = `				
				<pre>开奖条件满足，准备开奖</pre>
				<pre>中奖信息如下:%s</pre>				
				`
				msg := fmt.Sprintf(msgTpl, GItemList.RewordAll)
				fmt.Println(msg)
				//telegram.SendToBot(msg)
			}
			if findLimitN.IsLimitN == 2 {
				//第三种
				if GItemList.PayCount > 0 && gameOrderNum >= GItemList.PayCount {
					msgTpl = `				
					<pre>开奖条件满足，准备开奖</pre>
					<pre>中奖信息如下:%s</pre>				
					`
					msg := fmt.Sprintf(msgTpl, GItemList.RewordAll)
					fmt.Println(msg)
					//telegram.SendToBot(msg)
				}
			}

		}

		//fmt.Println(msgTpl)
		//msg := fmt.Sprintf(msgTpl, order.TradeId, order.OrderId, order.ActualAmount, order.Token, order.CreatedAt.ToDateTimeString(),carbon.Now().ToDateTimeString(), gameOrderNum, resNum)

	}

}

// // Trc20CallBack trc20回调
// func Trc20CallBack(token string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Sugar.Error(err)
// 		}
// 	}()
// 	client := http_client.GetHttpClient()
// 	startTime := carbon.Now().AddHours(-24).TimestampWithMillisecond()
// 	endTime := carbon.Now().TimestampWithMillisecond()
// 	resp, err := client.R().SetQueryParams(map[string]string{
// 		"sort":            "-timestamp",
// 		"reverse":         "true",
// 		"limit":           "50",
// 		"start":           "0",
// 		"direction":       "2",
// 		"db_version":      "1",
// 		"trc20Id":         "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
// 		"address":         token,
// 		"start_timestamp": stdutil.ToString(startTime),
// 		"end_timestamp":   stdutil.ToString(endTime),
// 	}).Get(UsdtTrc20ApiUri)

// 	if err != nil {
// 		panic(err)
// 	}
// 	if resp.StatusCode() != http.StatusOK {
// 		panic(err)
// 	}
// 	var trc20Resp UsdtTrc20Resp
// 	err = json.Cjson.Unmarshal(resp.Body(), &trc20Resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if trc20Resp.PageSize <= 0 {
// 		return
// 	}
// 	for _, transfer := range trc20Resp.Data {
// 		if transfer.To != token || transfer.ContractRet != "SUCCESS" {
// 			continue
// 		}
// 		decimalQuant, err := decimal.NewFromString(transfer.Amount)
// 		if err != nil {
// 			panic(err)
// 		}
// 		decimalDivisor := decimal.NewFromFloat(1000000)
// 		amount := decimalQuant.Div(decimalDivisor).InexactFloat64()
// 		tradeId, err := data.GetTradeIdByWalletAddressAndAmount(token, amount)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if tradeId == "" {
// 			continue
// 		}
// 		order, err := data.GetOrderInfoByTradeId(tradeId)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// 区块的确认时间必须在订单创建时间之后
// 		createTime := order.CreatedAt.TimestampWithMillisecond()
// 		if transfer.BlockTimestamp < createTime {
// 			panic("Orders cannot actually be matched")
// 		}
// 		// 到这一步就完全算是支付成功了
// 		req := &request.OrderProcessingRequest{
// 			Token:              token,
// 			TradeId:            tradeId,
// 			Amount:             amount,
// 			BlockTransactionId: transfer.Hash,
// 		}
// 		err = OrderProcessing(req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// 回调队列
// 		orderCallbackQueue, _ := handle.NewOrderCallbackQueue(order)
// 		mq.MClient.Enqueue(orderCallbackQueue, asynq.MaxRetry(5))
// 		// 发送机器人消息
// 		msgTpl := `
// <b>📢📢有新的交易支付成功！</b>
// <pre>交易号：%s</pre>
// <pre>订单号：%s</pre>
// <pre>请求支付金额：%f cny</pre>
// <pre>实际支付金额：%f usdt</pre>
// <pre>钱包地址：%s</pre>
// <pre>订单创建时间：%s</pre>
// <pre>支付成功时间：%s</pre>
// `
// 		msg := fmt.Sprintf(msgTpl, order.TradeId, order.OrderId, order.Amount, order.ActualAmount, order.Token, order.CreatedAt.ToDateTimeString(), carbon.Now().ToDateTimeString())
// 		telegram.SendToBot(msg)
// 	}
// }
