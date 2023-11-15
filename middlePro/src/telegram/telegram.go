package telegram

import (
	"time"

	"github.com/pengk/summer/config"
	"github.com/pengk/summer/model/data"
	"github.com/pengk/summer/util/log"
	tb "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var bots *tb.Bot

// BotStart 机器人启动
func BotStart() {
	var err error
	getGroupToken, _ := data.GetGroupBotSetting()
	// botSetting := tb.Settings{
	// 	Token:  config.TgBotToken,
	// 	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	// }
	botSetting := tb.Settings{
		Token:  getGroupToken.Tgtoken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}
	if config.TgProxy != "" {
		botSetting.URL = config.TgProxy
	}
	bots, err = tb.NewBot(botSetting)
	if err != nil {
		log.Sugar.Error(err.Error())
		return
	}
	err = bots.SetCommands(Cmds)
	if err != nil {
		log.Sugar.Error(err.Error())
		return
	}
	RegisterHandle()
	bots.Start()
}

// RegisterHandle 注册处理器
func RegisterHandle() {
	adminOnly := bots.Group()
	adminOnly.Use(middleware.Whitelist(config.TgManage))
	adminOnly.Handle(START_CMD, WalletList)
	bots.Handle(tb.OnText, OnTextSendMsgCodeHandle)
	//adminOnly.Handle(tb.OnText, OnTextMessageHandle)
}

// SendToBot 主动发送消息机器人消息
func SendToBot(msg string) {
	go func() {
		user := tb.User{
			ID: config.TgManage,
		}
		_, err := bots.Send(&user, msg, &tb.SendOptions{
			ParseMode: tb.ModeHTML,
		})
		if err != nil {
			log.Sugar.Error(err)
		}
	}()
}
