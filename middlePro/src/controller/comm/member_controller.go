package comm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/assimon/luuu/model/data"
	"github.com/assimon/luuu/model/mdb"
	"github.com/assimon/luuu/util/page"
	"github.com/labstack/echo/v4"
)

type CodeData struct {
	AppID   string `json:"appId"`
	Imel    string `json:"imel"`
	MPhone  string `json:"mphone"`
	RandStr string `json:"randStr"`
	Secret  string `json:"secrect"`
	TimeID  int64  `json:"timeId"`
}
type AuthCodeData struct {
	Authcode CodeData `json:"authcode"`
	Cuname   string   `json:"cuname"`
}
type WalletAddrssAdd struct {
	Cuid          int    `json:"cuid"`
	WalletAddress string `json:"wallet_address"`
}
type WithDrawReq struct {
	Cuid   int `json:"cuid"`
	Amount int `json:"amount"`
}

// 获取member
func (c *BaseCommController) GetMember(ctx echo.Context) (err error) {
	Mid := ctx.Param("tid")
	resp, err := data.GetMemById(Mid)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 修改用户钱包
func (c *BaseCommController) EditWalletAddress(ctx echo.Context) (err error) {
	requestData := new(WalletAddrssAdd)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	resp, err2 := data.EditMemberWalletAddress(requestData.Cuid, requestData.WalletAddress)
	if err != nil {
		return c.FailJson(ctx, err2)
	}
	return c.SucJson(ctx, resp)
}

// 获取token
func (c *BaseCommController) GetToken(ctx echo.Context) (err error) {
	requestData := new(AuthCodeData)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	var authcode = requestData.Authcode
	var json1Builder strings.Builder
	json1Builder.WriteString(authcode.AppID)
	json1Builder.WriteString(authcode.Secret)
	json1Builder.WriteString(authcode.Imel)
	json1Builder.WriteString(fmt.Sprintf("%d", authcode.TimeID))
	json1Builder.WriteString(authcode.MPhone)
	json1Builder.WriteString(authcode.RandStr)
	json1 := json1Builder.String()
	hasher := md5.New()
	hasher.Write([]byte(json1))
	json1MD5Bytes := hasher.Sum(nil)
	json2 := hex.EncodeToString(json1MD5Bytes)
	combinedData := json1 + json2
	hasher2 := md5.New()
	hasher2.Write([]byte(combinedData))
	md5Bytes := hasher2.Sum(nil)
	finalJSON := hex.EncodeToString(md5Bytes)
	return c.SucJson(ctx, finalJSON)
}

// 获得用户信息
func (c *BaseCommController) GetMemberInfo(ctx echo.Context) (err error) {
	cuid := ctx.Param("uid")
	fineMember, err := data.GetMemById(cuid)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, fineMember)
}

// 前端更改用户信息
func (c *BaseCommController) EditMemInfo(ctx echo.Context) (err error) {
	requestData := new(mdb.MemberModel)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	fineMember, err := data.EditMemInfoModel(requestData)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, fineMember)
}

func (c *BaseCommController) MakeWithdraw(ctx echo.Context) (err error) {
	requestData := new(WithDrawReq)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	Cuid := strconv.Itoa(requestData.Cuid)
	fineMember, err := data.GetMemById(Cuid)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	//插入withdraw 进账
	InsertWithdraw := &mdb.WithdrawModel{
		Cuid:          requestData.Cuid,
		Amount:        requestData.Amount,
		Ctype:         2,
		WalletAddress: fineMember.WalletAddress,
		Sta:           1,
	}
	totalU, _ := strconv.Atoi(fineMember.TotalU)
	if totalU >= requestData.Amount {
		data.SetInAccountResult(InsertWithdraw)
		//member 扣除U
		data.ReduceRewordTotalU(requestData.Amount, requestData.Cuid)
		fineMember, _ := data.GetMemById(Cuid)
		return c.SucJson(ctx, fineMember)
	} else {
		return c.FailSuccJson(ctx, "提现金额不足")
	}

}

//个人中心信息

func (c *BaseCommController) GetRewordNum(ctx echo.Context) (err error) {
	cuid := ctx.Param("uid")
	fineMember, err := data.GetRewordRecordCountByUid(cuid)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, fineMember)
}

// 中奖记录
func (c *BaseCommController) GetRewordLists(ctx echo.Context) (err error) {
	cuid := ctx.Param("uid")
	reqdata := new(page.PageInfo)
	fineMember, err := data.GetRewordListsByUid(cuid, reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, fineMember)
}

// 删除中奖记录
func (c *BaseCommController) DelRewordListsById(ctx echo.Context) (err error) {
	wid := ctx.Param("wid")
	return c.SucJson(ctx, wid)
}

// 提现记录
func (c *BaseCommController) GetWithdrawLists(ctx echo.Context) (err error) {
	cuid := ctx.Param("uid")
	reqdata := new(page.PageInfo)
	fineMember, err := data.GetWithDrawListsByUid(cuid, reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, fineMember)
}

// 删除提现记录
func (c *BaseCommController) DelWithdrawById(ctx echo.Context) (err error) {
	wid := ctx.Param("wid")
	data.DeleteWithDrawSta(wid)
	return c.SucJson(ctx, wid)
}
