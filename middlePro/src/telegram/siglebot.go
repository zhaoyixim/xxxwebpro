package telegram

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/assimon/luuu/config"
	"github.com/assimon/luuu/model/data"
	"github.com/assimon/luuu/model/mdb"
	"github.com/assimon/luuu/util/log"
	tb "gopkg.in/telebot.v3"
)

var bot *tb.Bot
var subscriptions = make(map[int][]string)

func SigleBot() {
	var err error
	getSignleToken, _ := data.GetSigleBotSetting()
	// botSetting := tb.Settings{
	// 	Token:  config.TgBotToken,
	// 	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	// }
	botSetting := tb.Settings{
		Token:  getSignleToken.Tgtoken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	if config.TgProxy != "" {
		botSetting.URL = config.TgProxy
	}
	fmt.Println(botSetting)
	bot, err = tb.NewBot(botSetting)
	fmt.Println(err)
	if err != nil {
		fmt.Println(bot)
		log.Sugar.Error(err.Error())
		return
	}

	bot.Handle("/start", func(c tb.Context) error {
		user := c.Sender()
		// 发送欢迎消息
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(900000) + 100001 // 生成 100001 到 999999 之间的随机数
		randomNum := strconv.Itoa(randomNumber)
		message := "您的验证码是" + strconv.Itoa(randomNumber) + "请在10分钟内完成验证，过期将需要重新获取"
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
		_, err := bot.Send(user, message)
		if err != nil {
			log.Sugar.Error(err.Error())
		}
		return nil
	})

	// 处理私聊消息
	bot.Handle(tb.OnText, func(ctx tb.Context) error {
		if ctx.Message().Private() {
			reply := "这是一个私聊消息。"
			fmt.Println("这是一个私聊消息。")
			_, err := bot.Send(ctx.Message().Chat, reply)
			if err != nil {
				log.Sugar.Error(err.Error())
			}
		}
		return nil
	})

	bot.Start()
}
