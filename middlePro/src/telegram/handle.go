package telegram

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
	"github.com/pengk/summer/model/data"
	"github.com/pengk/summer/model/mdb"
	"github.com/pengk/summer/util/log"
	tb "gopkg.in/telebot.v3"
)

const (
	ReplayAddWallet = "请发给我一个合法的钱包地址"
	SendRegMsgCode  = "发送验证码"
)

func OnTextSendMsgCodeHandle(c tb.Context) error {
	if c.Message().ReplyTo.Text == SendRegMsgCode {
		defer bots.Delete(c.Message().ReplyTo)

		user := c.Sender()
		// 发送欢迎消息
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(900000) + 100001 // 生成 100001 到 999999 之间的随机数
		randomNum := strconv.Itoa(randomNumber)
		isbot := 0
		if user.IsBot {
			isbot = 1
		}
		botdatar := mdb.MsgCode{
			Msgcode:     randomNum,
			TgId:        user.ID,
			TgFirstname: user.FirstName,
			TgLastname:  user.LastName,
			TgUsername:  user.Username,
			TgIsbot:     isbot,
		}
		errs := data.GenerateMsgCode(&botdatar)

		if errs != nil {
			log.Sugar.Error(errs)
		}
		c.Send(fmt.Sprintf("您的验证码是[%s],请在10分钟内完成验证，过期将需要重新获取！", randomNum))
		return WalletList(c)
	}
	return nil
}

func OnTextMessageHandle(c tb.Context) error {
	if c.Message().ReplyTo.Text == ReplayAddWallet {
		defer bots.Delete(c.Message().ReplyTo)
		_, err := data.AddWalletAddress(c.Message().Text)
		if err != nil {
			return c.Send(err.Error())
		}
		c.Send(fmt.Sprintf("钱包[%s]添加成功！", c.Message().Text))
		return WalletList(c)
	}
	return nil
}

func WalletList(c tb.Context) error {
	wallets, err := data.GetAllWalletAddress()
	if err != nil {
		return err
	}
	var btnList [][]tb.InlineButton
	for _, wallet := range wallets {
		status := "已启用✅"
		if wallet.Status == mdb.TokenStatusDisable {
			status = "已禁用🚫"
		}
		var temp []tb.InlineButton
		btnInfo := tb.InlineButton{
			Unique: wallet.Token,
			Text:   fmt.Sprintf("%s[%s]", wallet.Token, status),
			Data:   strutil.MustString(wallet.ID),
		}
		bots.Handle(&btnInfo, WalletInfo)
		btnList = append(btnList, append(temp, btnInfo))
	}
	addBtn := tb.InlineButton{Text: "添加钱包地址", Unique: "AddWallet"}
	bots.Handle(&addBtn, func(c tb.Context) error {
		return c.Send(ReplayAddWallet, &tb.ReplyMarkup{
			ForceReply: true,
		})
	})
	btnList = append(btnList, []tb.InlineButton{addBtn})
	return c.EditOrSend("请点击钱包继续操作", &tb.ReplyMarkup{
		InlineKeyboard: btnList,
	})
}

func WalletInfo(c tb.Context) error {
	id := mathutil.MustUint(c.Data())
	tokenInfo, err := data.GetWalletAddressById(id)
	if err != nil {
		return c.Send(err.Error())
	}
	enableBtn := tb.InlineButton{
		Text:   "启用",
		Unique: "enableBtn",
		Data:   c.Data(),
	}
	disableBtn := tb.InlineButton{
		Text:   "禁用",
		Unique: "disableBtn",
		Data:   c.Data(),
	}
	delBtn := tb.InlineButton{
		Text:   "删除",
		Unique: "delBtn",
		Data:   c.Data(),
	}
	backBtn := tb.InlineButton{
		Text:   "返回",
		Unique: "WalletList",
	}
	bots.Handle(&enableBtn, EnableWallet)
	bots.Handle(&disableBtn, DisableWallet)
	bots.Handle(&delBtn, DelWallet)
	bots.Handle(&backBtn, WalletList)
	return c.EditOrReply(tokenInfo.Token, &tb.ReplyMarkup{InlineKeyboard: [][]tb.InlineButton{
		{
			enableBtn,
			disableBtn,
			delBtn,
		},
		{
			backBtn,
		},
	}})
}

func EnableWallet(c tb.Context) error {
	id := mathutil.MustUint(c.Data())
	if id <= 0 {
		return c.Send("请求不合法！")
	}
	err := data.ChangeWalletAddressStatus(id, mdb.TokenStatusEnable)
	if err != nil {
		return c.Send(err.Error())
	}
	return WalletList(c)
}

func DisableWallet(c tb.Context) error {
	id := mathutil.MustUint(c.Data())
	if id <= 0 {
		return c.Send("请求不合法！")
	}
	err := data.ChangeWalletAddressStatus(id, mdb.TokenStatusDisable)
	if err != nil {
		return c.Send(err.Error())
	}
	return WalletList(c)
}

func DelWallet(c tb.Context) error {
	id := mathutil.MustUint(c.Data())
	if id <= 0 {
		return c.Send("请求不合法！")
	}
	err := data.DeleteWalletAddressById(id)
	if err != nil {
		return c.Send(err.Error())
	}
	return WalletList(c)
}
