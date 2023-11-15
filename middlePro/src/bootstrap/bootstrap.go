package bootstrap

import (
	"github.com/pengk/summer/command"
	"github.com/pengk/summer/config"
	"github.com/pengk/summer/model/dao"
	"github.com/pengk/summer/mq"
	"github.com/pengk/summer/task"
	"github.com/pengk/summer/telegram"
	"github.com/pengk/summer/util/log"
)

// Start 服务启动
func Start() {
	// 配置加载
	config.Init()
	// 日志加载
	log.Init()
	// Mysql启动
	dao.MysqlInit()
	// redis启动
	dao.RedisInit()
	// 队列启动
	mq.Start()
	// telegram机器人启动
	go telegram.BotStart()
	//go telegram.SigleBot()
	// 定时任务
	go task.Start()
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}
