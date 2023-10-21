package comm

import (
	"github.com/assimon/luuu/model/data"
	"github.com/assimon/luuu/model/request"
	"github.com/assimon/luuu/model/response"
	"github.com/assimon/luuu/model/service"
	"github.com/assimon/luuu/util/constant"

	"github.com/labstack/echo/v4"
)

// CreateTransaction 创建交易
func (c *BaseCommController) CreateTransaction(ctx echo.Context) (err error) {
	req := new(request.CreateTransactionRequest)
	if err = ctx.Bind(req); err != nil {
		return c.FailJson(ctx, constant.ParamsMarshalErr)
	}
	if err = c.ValidateStruct(ctx, req); err != nil {
		return c.FailJson(ctx, err)
	}
	resp, err := service.CreateTransaction(req)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 根据itemid 提取成功付款的人数
func (c *BaseCommController) GetPayedOrdered(ctx echo.Context) (err error) {
	itemid := ctx.Param("itemid")
	findItem, _ := data.GetGDtlById(itemid)
	respOrderItems, _ := data.GetOrderInfoByItemId(itemid)
	findLimitN, _ := data.GetLimitNumByGListId(itemid, findItem.Cuid)
	targetOrderLen := len(respOrderItems.Corders)
	totallPayedU := targetOrderLen * findItem.PayUnit
	if findLimitN.IsLimitN == 0 {
		//第一种
		if totallPayedU >= findItem.TotalU || findItem.Sta == 2 {
			resp := response.ItemTimesReturns{
				Message:     "该项目已经筹款结束，等待发布中奖名单",
				Status:      2,
				RestTimes:   0,
				OrdersData:  respOrderItems.Corders,
				MembersData: respOrderItems.Members,
				ItemSta:     findItem.Sta,
			}
			return c.SucJson(ctx, resp)
		}
	}
	//第二种不用管-手动停止
	if findLimitN.IsLimitN == 1 {
		if findItem.Sta == 2 {
			resp := response.ItemTimesReturns{
				Message:     "游戏结束",
				Status:      2,
				RestTimes:   0,
				OrdersData:  respOrderItems.Corders,
				MembersData: respOrderItems.Members,
				ItemSta:     findItem.Sta,
			}
			return c.SucJson(ctx, resp)
		}

	}
	if findLimitN.IsLimitN == 2 {
		//第三种
		if targetOrderLen >= findItem.PayCount || findItem.Sta == 2 {
			resp := response.ItemTimesReturns{
				Message:     "设置的筹款总笔数已经用完",
				Status:      2,
				RestTimes:   0,
				OrdersData:  respOrderItems.Corders,
				MembersData: respOrderItems.Members,
				ItemSta:     findItem.Sta,
			}
			return c.SucJson(ctx, resp)
		}
	}

	//reword_count
	resttimes := findItem.RewordCount - len(respOrderItems.Corders)
	resp := response.ItemTimesReturns{
		Message:     "",
		Status:      1,
		RestTimes:   resttimes,
		OrdersData:  respOrderItems.Corders,
		MembersData: respOrderItems.Members,
		ItemSta:     findItem.Sta,
	}

	return c.SucJson(ctx, resp)
}
