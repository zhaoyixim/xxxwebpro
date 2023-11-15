package route

import (
	"github.com/labstack/echo/v4"
	cmiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/pengk/summer/controller/comm"
	v1 "github.com/pengk/summer/controller/comm/v1"
	"github.com/pengk/summer/middleware"
)

// RegisterRoute 路由注册
func RegisterRoute(e *echo.Echo) {
	// e.Any("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "hello pengk, https://github.com/assimon/pengk")
	// })
	e.Static("/upload", "upload")

	/*添加跨域访问*/
	//e.Use(cmiddleware.CORS())
	e.Use(cmiddleware.CORSWithConfig(cmiddleware.CORSConfig{
		AllowOrigins: []string{"*"}, // 允许所有域名跨域访问，可以替换为你的域名
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		//AllowHeaders: []string{echo.HeaderOrigin, "timestamp", echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, // 允许的头部字段
	}))
	//game 相关
	gameRoute := e.Group("/g")
	//获取token
	gameRoute.POST("/gettoken", comm.Ctrl.GetToken)
	//登录
	gameRoute.POST("/login", comm.Ctrl.Login)
	//设置过期
	gameRoute.GET("/setexpire/:uid", comm.Ctrl.Login)
	//获取首页相关信息
	gameRoute.GET("/getlists", comm.Ctrl.IndexGLists)
	//获取我能够发布的游戏列表
	gameRoute.GET("/getlist/:uid", comm.Ctrl.IndexGMyList)
	//检查是否有权限发布游戏
	gameRoute.POST("/checkispub", comm.Ctrl.CheckIsPublish)
	//获取轮播图
	gameRoute.GET("/getswiper", comm.Ctrl.IndexGSwiper)
	//中间菜单列表----进行中的所有列表
	gameRoute.POST("/getallunfinishedlists", comm.Ctrl.GetAllDodingG)
	//通过ID查询列表-进行中
	gameRoute.POST("/getunfinishedlist/:tid", comm.Ctrl.GetUnfinishedGList)
	//通过ID查询列表-已完成
	gameRoute.POST("/getfinishedlist/:tid", comm.Ctrl.GetfinishedGList)
	//游戏发布记录
	//完成记录
	gameRoute.POST("/publishfinishedlist/:uid", comm.Ctrl.GetPublishfinished)
	//未完成记录
	gameRoute.POST("/publishunfinishedlist/:uid", comm.Ctrl.GetPublishUnfinished)
	//游戏详情
	gameRoute.POST("/getdtl", comm.Ctrl.GetGDtl)
	//发布游戏的列表
	gameRoute.GET("/publist", comm.Ctrl.GetAllDodingG)
	//发布游戏
	gameRoute.POST("/publish", comm.Ctrl.PubGame, middleware.AuthorizationMiddleware)
	//更新用户地址钱包
	gameRoute.POST("/editwalletaddress", comm.Ctrl.EditWalletAddress)
	//获取用户信息
	gameRoute.GET("/memberinfo/:uid", comm.Ctrl.GetMemberInfo)
	//前端更改用户信息
	gameRoute.POST("/changememinfo", comm.Ctrl.EditMemInfo)
	//提现申请
	gameRoute.POST("/withdraw", comm.Ctrl.MakeWithdraw)
	//提现订单fmt.Println(botSetting)
	//gameRoute.GET("/getdetail/:id", comm.Ctrl.IndexGLists)
	// 用户个人中心的信息
	gameRoute.POST("/getrewordcount/:uid", comm.Ctrl.GetRewordNum)
	//中奖记录
	gameRoute.POST("/getrewordlist/:uid", comm.Ctrl.GetRewordLists)
	gameRoute.POST("/delreword/:wid", comm.Ctrl.DelRewordListsById)

	//提现记录
	gameRoute.POST("/getwithdrawlist/:uid", comm.Ctrl.GetWithdrawLists)
	gameRoute.POST("/delwithdraw/:wid", comm.Ctrl.DelWithdrawById)

	//前台获取配置
	gameRoute.POST("/getconfigbyname", comm.Ctrl.GetConfigByName)

	// ==== 支付相关=====
	payRoute := e.Group("/pay")
	// 收银台
	payRoute.GET("/checkout-counter/:trade_id", comm.Ctrl.CheckoutCounter)
	// 状态检测
	payRoute.GET("/check-status/:trade_id", comm.Ctrl.CheckStatus)

	//通过itemid获取已支付的订单
	payRoute.GET("/getpayedinfo/:itemid", comm.Ctrl.GetPayedOrdered)
	apiV1Route := e.Group("/api/v1")

	// ====订单相关====
	orderRoute := apiV1Route.Group("/order", middleware.CheckApiSign())
	// 创建订单
	orderRoute.POST("/create-transaction", comm.Ctrl.CreateTransaction)

	/*后台操作部分*/
	apiV1Route.POST("/userlogin", v1.Ctrl.Login)
	apiV1Route.GET("/useradmininfo", v1.Ctrl.AdminInfo)
	apiV1Route.GET("/dashboardconsole", v1.Ctrl.DashboardConsole)
	apiV1Route.POST("/tablewithdrawList", v1.Ctrl.WithdrawLists)
	apiV1Route.POST("/sureWithDraw", v1.Ctrl.SureWithDraw)
	apiV1Route.POST("/depositList", v1.Ctrl.DepositList)
	apiV1Route.POST("/userlist", v1.Ctrl.BGetUserLists)
	apiV1Route.POST("/gamegetLists", v1.Ctrl.BGetGameLists)
	apiV1Route.POST("/sysgetwebconfig", v1.Ctrl.SysGetWebConfig)
	apiV1Route.POST("/sysaddwebconfig", v1.Ctrl.PSysAddWebConfig)
	apiV1Route.POST("/updatesysconfig", v1.Ctrl.UpdateWebConfig)
	apiV1Route.POST("/updatewebconfigopen", v1.Ctrl.UpdateWebConfigOpen)
	apiV1Route.POST("/updatewebconfigclose", v1.Ctrl.UpdateWebConfigClose)

	//类别列表
	apiV1Route.POST("/gameitemlists", v1.Ctrl.BGetGameItemLists)
	apiV1Route.POST("/additemdata", v1.Ctrl.AddItemData)
	apiV1Route.POST("/editclassitem", v1.Ctrl.EditItemData)
	apiV1Route.POST("/deleteclassitem", v1.Ctrl.DeleteItemData)

	//确定商户身份
	apiV1Route.POST("/usersuremech", v1.Ctrl.SureMechId)
	apiV1Route.POST("/operashowitem", v1.Ctrl.OperaItemShow)

	//获得收款钱包
	apiV1Route.POST("/getwalletlists", v1.Ctrl.GetWalletLists)
	apiV1Route.POST("/createwalletandmech", v1.Ctrl.CreateMechWalllet)
	apiV1Route.POST("/forbidwallet", v1.Ctrl.ForbidWalletMech)

	//获取匹配数据
	apiV1Route.POST("/getmatchitems", v1.Ctrl.GetMatchItems)
	apiV1Route.POST("/additemmatchdata", v1.Ctrl.AddItemMatchData)
	apiV1Route.POST("/edititemmatchdata", v1.Ctrl.EditItemMatchData)
	apiV1Route.POST("/matchupdateopen", v1.Ctrl.EditItemMatchOpen)
	apiV1Route.POST("/matchupdateclose", v1.Ctrl.EditItemMatchClose)
	//机器人设置部分
	apiV1Route.POST("/getbotsettings", v1.Ctrl.GetBotSettings)
	//添加数据
	apiV1Route.POST("/addbotsetting", v1.Ctrl.AddBotSetting)
	apiV1Route.POST("/editbotsetting", v1.Ctrl.EditBotSetting)
	apiV1Route.POST("/openorclosebotset", v1.Ctrl.OpenorCloseBotSet)
	//轮播图
	apiV1Route.GET("/sysgetImages", v1.Ctrl.SysGetImages)
	apiV1Route.POST("/apimultiupload", v1.Ctrl.ApiMultiUpload)
	apiV1Route.POST("/deleteImg", v1.Ctrl.DeleteImg)

}
