package comm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pengk/summer/model/data"
	"github.com/pengk/summer/model/mdb"
)

type CreateLogin struct {
	Cuname  string `json:"cuname" validate:"required|maxLen:32"`
	Passwd  string `json:"passwd" validate:"required"`
	IsLogin int    `json:"is_login" validate:"required"`
	Msgcode string `json:"msgcode"`
}

// 登录
func (c *BaseCommController) Login(ctx echo.Context) (err error) {
	reqdata := new(CreateLogin)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	//查询msgcode
	if reqdata.IsLogin == 0 {
		findCodeMsg, codeerr := data.FindMsgCodeByCode(reqdata.Msgcode)
		if codeerr != nil {
			return c.FailJson(ctx, codeerr)
		}
		if findCodeMsg == nil {
			returnerr := "请输入验证码"
			return c.FailSuccJson(ctx, returnerr)
		}
	}
	ip := ctx.RealIP()
	hash := md5.Sum([]byte(reqdata.Passwd))
	passwdhashHex := hex.EncodeToString(hash[:])
	findUsername, _ := data.FindByUseranme(reqdata.Cuname)
	reJson := "登录成功"

	//登录
	if reqdata.IsLogin == 1 {
		if findUsername.ID > 0 {

			if passwdhashHex != findUsername.Passwd {
				reJson = "密码错误"
				return c.FailSuccJson(ctx, reJson)
			}
			memErr := data.LoginAndUpdata(findUsername.ID)
			if memErr != nil {
				return c.FailJson(ctx, memErr)
			}

			return c.SucJson(ctx, findUsername, reJson)
		} else {
			reJson = "您还未注册"
			return c.FailSuccJson(ctx, reJson)
		}

	} else {
		//注册
		if findUsername.ID > 0 {
			returnerr := "已有用户，请登录,有任何问题请联系客服"
			return c.FailSuccJson(ctx, returnerr)
		}
		//reqmsgcode, _ := strconv.Atoi(reqdata.Msgcode)

		if reqdata.Msgcode == "" {
			returnerr := "请输入验证码"
			return c.FailSuccJson(ctx, returnerr)
		}

		fmt.Println(reqdata)
		//查看是否点击产生Msgcode
		findmsg, msgerr := data.FindMsgCodeByCode(reqdata.Msgcode)
		if findmsg.ID == 0 {
			returnerr := "验证码无效"
			return c.FailSuccJson(ctx, returnerr)
		}
		if msgerr != nil || findmsg.Sta != 0 {
			returnerr := "验证码失效"
			return c.FailSuccJson(ctx, returnerr)
		}
		//添加数据
		AddMemberData := mdb.MemberModel{
			Cuname:  reqdata.Cuname,
			Ip:      ip,
			Msgcode: reqdata.Msgcode,
			Passwd:  passwdhashHex,
		}
		memErr := data.RegAndAddmember(&AddMemberData)
		if memErr != nil {
			return c.FailJson(ctx, memErr)
		}
		findUsername, _ := data.FindByUseranme(reqdata.Cuname)
		//更新msg 为已使用
		_, _ = data.UpdataMsgCodeByCode(reqdata.Msgcode)
		reJson = "注册成功"
		return c.SucJson(ctx, findUsername, reJson)
	}

	// fmt.Println(resp)

	// tmpl, err := template.ParseFiles(fmt.Sprintf(".%s/%s", config.StaticPath, "index.html"))
	// if err != nil {
	// 	return ctx.String(http.StatusOK, err.Error())
	// }
	// return tmpl.Execute(ctx.Response(), resp)
}
