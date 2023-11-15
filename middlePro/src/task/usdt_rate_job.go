package task

import (
	"time"

	"github.com/pengk/summer/config"
	"github.com/pengk/summer/util/http_client"
	"github.com/pengk/summer/util/json"
	"github.com/pengk/summer/util/log"
	"github.com/pengk/summer/util/math"
)

const UsdtRateApiUri = "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/detail/chart"

type UsdtRateJob struct {
}

type UsdtRateResp struct {
	Data   Data   `json:"data"`
	Status Status `json:"status"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    string    `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      string    `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}

type Data struct {
	Points map[string]Points `json:"points"`
}

type Points struct {
	V []float64 `json:"v"`
	C []float64 `json:"c"`
}

func (r UsdtRateJob) Run() {
	client := http_client.GetHttpClient()
	resp, err := client.R().SetQueryString("id=825&range=1H&convertId=2787").SetHeader("Accept", "application/json").Get(UsdtRateApiUri)
	if err != nil {
		log.Sugar.Error("usdt rate get err:", err.Error())
		return
	}
	var usdtResp UsdtRateResp
	err = json.Cjson.Unmarshal(resp.Body(), &usdtResp)
	if err != nil {
		log.Sugar.Error("Unmarshal usdt resp err:", err.Error())
		return
	}
	if usdtResp.Status.ErrorCode != "0" {
		log.Sugar.Error("usdt resp err:", usdtResp.Status.ErrorMessage)
		return
	}
	for _, points := range usdtResp.Data.Points {
		if len(points.C) > 0 && points.C[0] > 0 {
			config.UsdtRate = math.MustParsePrecFloat64(points.C[0], 2)
			return
		}
	}
}
