package comm

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pengk/summer/config"
	"github.com/pengk/summer/model/response"
	"github.com/pengk/summer/model/service"
)

// 新收银台
func (c *BaseCommController) CheckoutPay(ctx echo.Context) (err error) {
	tradeId := ctx.Param("trade_id")
	resp, err := service.GetCheckoutCounterByTradeId(tradeId)
	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}
	tmpl, err := template.ParseFiles(fmt.Sprintf(".%s/%s", config.StaticPath, "index.html"))
	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}
	return tmpl.Execute(ctx.Response(), resp)
}

// CheckoutCounter 收银台
func (c *BaseCommController) CheckoutCounter(ctx echo.Context) (err error) {
	tradeId := ctx.Param("trade_id")
	resp, err := service.GetCheckoutCounterByTradeId(tradeId)
	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}
	tmpl, err := template.ParseFiles(fmt.Sprintf(".%s/%s", config.StaticPath, "index.html"))
	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}
	return tmpl.Execute(ctx.Response(), resp)
}

// CheckStatus 支付状态检测
func (c *BaseCommController) CheckStatus(ctx echo.Context) (err error) {
	tradeId := ctx.Param("trade_id")
	order, err := service.GetOrderInfoByTradeId(tradeId)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	resp := response.CheckStatusResponse{
		TradeId: order.TradeId,
		Status:  order.Status,
	}
	return c.SucJson(ctx, resp)
}
