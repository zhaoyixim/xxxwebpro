package handle

import (
	"context"
	"errors"

	"github.com/hibiken/asynq"
	"github.com/pengk/summer/config"
	"github.com/pengk/summer/model/data"
	"github.com/pengk/summer/model/mdb"
	"github.com/pengk/summer/model/response"
	"github.com/pengk/summer/util/http_client"
	"github.com/pengk/summer/util/json"
	"github.com/pengk/summer/util/log"
	"github.com/pengk/summer/util/sign"
)

const QueueOrderCallback = "order:callback"

func NewOrderCallbackQueue(order *mdb.Orders) (*asynq.Task, error) {
	payload, err := json.Cjson.Marshal(order)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(QueueOrderCallback, payload), nil
}

func OrderCallbackHandle(ctx context.Context, t *asynq.Task) error {
	var order mdb.Orders
	err := json.Cjson.Unmarshal(t.Payload(), &order)
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			log.Sugar.Error(err)
		}
	}()
	defer func() {
		data.SaveCallBackOrdersResp(&order)
	}()
	client := http_client.GetHttpClient()
	orderResp := response.OrderNotifyResponse{
		TradeId:            order.TradeId,
		OrderId:            order.OrderId,
		Amount:             order.Amount,
		ActualAmount:       order.ActualAmount,
		Token:              order.Token,
		BlockTransactionId: order.BlockTransactionId,
		Status:             mdb.StatusPaySuccess,
	}
	signature, err := sign.Get(orderResp, config.GetApiAuthToken())
	if err != nil {
		return err
	}
	orderResp.Signature = signature
	resp, err := client.R().SetHeader("powered-by", "pengk(https://github.com/assimon/pengk)").SetBody(orderResp).Post(order.NotifyUrl)
	if err != nil {
		return err
	}
	body := string(resp.Body())
	if body != "ok" {
		order.CallBackConfirm = mdb.CallBackConfirmNo
		return errors.New("not ok")
	}
	order.CallBackConfirm = mdb.CallBackConfirmOk
	return nil
}
