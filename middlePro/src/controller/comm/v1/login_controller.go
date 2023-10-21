package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/assimon/luuu/model/data"
	"github.com/assimon/luuu/model/mdb"

	"github.com/assimon/luuu/controller"
	"github.com/assimon/luuu/util/page"
	"github.com/labstack/echo/v4"
)

var Ctrl = &BaseCommController{}

type BaseCommController struct {
	controller.BaseController
}
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type TokenResponse struct {
	Token string `json:"token"`
}
type ConsoleInfo struct {
	Visits struct {
		DayVisits float64 `json:"dayVisits"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"visits"`
	Saleroom struct {
		WeekSaleroom float64 `json:"weekSaleroom"`
		Amount       float64 `json:"amount"`
		Degree       float64 `json:"degree"`
	} `json:"saleroom"`
	OrderLarge struct {
		WeekLarge float64 `json:"weekLarge"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"orderLarge"`
	Volume struct {
		WeekLarge float64 `json:"weekLarge"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"volume"`
}

func (c *BaseCommController) Login(ctx echo.Context) (err error) {
	data := TokenResponse{
		Token: "affsdfasdfa",
	}
	return c.BSucJson(ctx, data)
}

func (c *BaseCommController) AdminInfo(ctx echo.Context) (err error) {
	jsonStr := `{"userId":"1","username":"admin","realName":"Admin","avatar":"Random.image()","desc":"manager","password":"Random.string('upper',4,16)","token":"","permissions":[{"label":"主控台","value":"dashboard_console"},{"label":"监控页","value":"dashboard_monitor"},{"label":"工作台","value":"dashboard_workplace"},{"label":"基础列表","value":"basic_list"},{"label":"基础列表删除","value":"basic_list_delete"}]}`
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)
	return c.BSucJson(ctx, data)
}
func (c *BaseCommController) DashboardConsole(ctx echo.Context) (err error) {
	rand.Seed(time.Now().UnixNano())
	Visits, _ := data.BackGetRegVisits()
	consoleInfo := ConsoleInfo{
		Visits: struct {
			DayVisits float64 `json:"dayVisits"`
			Rise      float64 `json:"rise"`
			Decline   float64 `json:"decline"`
			Amount    float64 `json:"amount"`
		}{
			DayVisits: Visits.AllVisits,
			Rise:      Visits.DayVisits,
			Decline:   rand.Float64()*(99-10) + 10,
			Amount:    Visits.DayActive,
		},
		Saleroom: struct {
			WeekSaleroom float64 `json:"weekSaleroom"`
			Amount       float64 `json:"amount"`
			Degree       float64 `json:"degree"`
		}{
			WeekSaleroom: rand.Float64()*(99999-10000) + 10000,
			Amount:       rand.Float64()*(999999-99999) + 99999,
			Degree:       rand.Float64()*(99-10) + 10,
		},
		OrderLarge: struct {
			WeekLarge float64 `json:"weekLarge"`
			Rise      float64 `json:"rise"`
			Decline   float64 `json:"decline"`
			Amount    float64 `json:"amount"`
		}{
			WeekLarge: rand.Float64()*(99999-10000) + 10000,
			Rise:      rand.Float64()*(99-10) + 10,
			Decline:   rand.Float64()*(99-10) + 10,
			Amount:    rand.Float64()*(999999-99999) + 99999,
		},
		Volume: struct {
			WeekLarge float64 `json:"weekLarge"`
			Rise      float64 `json:"rise"`
			Decline   float64 `json:"decline"`
			Amount    float64 `json:"amount"`
		}{
			WeekLarge: rand.Float64()*(99999-10000) + 10000,
			Rise:      rand.Float64()*(99-10) + 10,
			Decline:   rand.Float64()*(99-10) + 10,
			Amount:    rand.Float64()*(999999-99999) + 99999,
		},
	}
	return c.BSucJson(ctx, consoleInfo)
}

// 提现列表
func (c *BaseCommController) WithdrawLists(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetWithdrawOut(pagedata)
	count, _ := data.GetWithdrawOutCount()
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

type SureInData struct {
	WithDrawId int `json:"withDrawId"`
}

// SureWithDraw
func (c *BaseCommController) SureWithDraw(ctx echo.Context) (err error) {
	reqdata := new(SureInData)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.UpdataWithDrawSta(reqdata.WithDrawId)
	return c.BSucJson(ctx, reqdata)
}
func (c *BaseCommController) DepositList(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetWithdrawIn(pagedata)
	count, _ := data.GetWithdrawInCount()
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}
func (c *BaseCommController) TorderList(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetRewordRecordList(pagedata)
	count, _ := data.GetRewordRecordCount()
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

/*后台*/
//用户信息
func (c *BaseCommController) BGetUserLists(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.BGetAllUserLists(pagedata)
	count, _ := data.BGetAllUserCount()
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}
func (c *BaseCommController) BGetGameLists(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetPublishLists(pagedata)
	count, _ := data.GetPublishListsCount()
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

func (c *BaseCommController) SysGetWebConfig(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetConfigLists(pagedata)
	count, _ := data.GetConfigListsCount()
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

type AddConfigReq struct {
	Cname  string `json:"cname"`
	Ckey   string `json:"ckey"`
	Cvalue string `json:"cvalue"`
	Mechid int    `json:"mech_id"`
}

func (c *BaseCommController) PSysAddWebConfig(ctx echo.Context) (err error) {
	requestData := new(mdb.SysConfig)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	err2 := data.SysAddWebConfig(requestData)
	if err != nil {
		return c.FailJson(ctx, err2)
	}
	return c.BSucJson(ctx, requestData)
}

// 更新配置
func (c *BaseCommController) UpdateWebConfig(ctx echo.Context) (err error) {
	requestData := new(mdb.SysConfig)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	err2 := data.SysUpdateConfig(requestData)
	if err != nil {
		return c.FailJson(ctx, err2)
	}
	return c.BSucJson(ctx, requestData)
}

// 更新配置-启用
func (c *BaseCommController) UpdateWebConfigOpen(ctx echo.Context) (err error) {
	requestData := new(mdb.SysConfig)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	err2 := data.SysUpdateConfigOpen(requestData)
	if err != nil {
		return c.FailJson(ctx, err2)
	}
	return c.BSucJson(ctx, requestData)
}

// 更新配置-禁用
func (c *BaseCommController) UpdateWebConfigClose(ctx echo.Context) (err error) {
	requestData := new(mdb.SysConfig)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	err2 := data.SysUpdateConfigClose(requestData)
	if err != nil {
		return c.FailJson(ctx, err2)
	}
	return c.BSucJson(ctx, requestData)
}

// 类别列表
func (c *BaseCommController) BGetGameItemLists(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetItemClassLists(pagedata)
	count, _ := data.GetItemClassListsCount()
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

// 确定商户身份
type SureMechId struct {
	Cuid int `json:"cuid"`
}

func (c *BaseCommController) SureMechId(ctx echo.Context) (err error) {
	reqdata := new(SureMechId)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.ChangeMemMechId(reqdata.Cuid)
	return c.BSucJson(ctx, reqdata)
}

type OperaItem struct {
	Mid int `json:"mid"`
}

// 操作类别展示
func (c *BaseCommController) OperaItemShow(ctx echo.Context) (err error) {
	reqdata := new(OperaItem)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.OperaItemShow(reqdata.Mid)
	return c.BSucJson(ctx, reqdata)
}

// 添加类别
func (c *BaseCommController) AddItemData(ctx echo.Context) (err error) {
	reqdata := new(mdb.GameItems)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.AddSysItemsData(reqdata)
	return c.BSucJson(ctx, reqdata)
}

// 编辑类别
func (c *BaseCommController) EditItemData(ctx echo.Context) (err error) {
	reqdata := new(mdb.GameItems)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	fmt.Println(reqdata)
	data.EditSysItemsData(reqdata)
	return c.BSucJson(ctx, reqdata)
}

//删除

type ClassItemId struct {
	ID int ` json:"id"`
}

func (c *BaseCommController) DeleteItemData(ctx echo.Context) (err error) {
	reqdata := new(ClassItemId)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	fmt.Println(reqdata)
	data.DeleteSysItemsData(reqdata.ID)
	return c.BSucJson(ctx, reqdata)
}

// 获得收款钱包
func (c *BaseCommController) GetWalletLists(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetAllWalletAddress()
	count := len(resp)
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

// AddWalletAddress 创建钱包
func (c *BaseCommController) CreateMechWalllet(ctx echo.Context) (err error) {
	reqdata := new(mdb.WalletAddress)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	address, _ := data.AddWalletAddressAndMechId(reqdata)
	respJson := page.BPageResponse{
		PageCount: 1,
		Page:      1,
		PageSize:  10,
		List:      address,
	}
	return c.BSucJson(ctx, respJson)
}

// 禁用钱包
type ForbidWid struct {
	Wid int `json:"wid"`
}

func (c *BaseCommController) ForbidWalletMech(ctx echo.Context) (err error) {
	reqdata := new(ForbidWid)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.ForBidWalletAddress(reqdata.Wid)
	return c.BSucJson(ctx, reqdata)
}

// 获取匹配数据
func (c *BaseCommController) GetMatchItems(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetAllMatchData()
	count := len(resp)
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

// 添加匹配数据
func (c *BaseCommController) AddItemMatchData(ctx echo.Context) (err error) {
	reqdata := new(mdb.CuidItemModel)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.AddItemMatchData(reqdata)
	return c.BSucJson(ctx, reqdata)
}
func (c *BaseCommController) EditItemMatchData(ctx echo.Context) (err error) {
	reqdata := new(mdb.CuidItemModel)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.EditItemMatchData(reqdata)
	return c.BSucJson(ctx, reqdata)
}

// 启用
// 获取配置
type ItemMatchReq struct {
	Id   int `json:"id"`
	Cuid int `json:"cuid"`
}

func (c *BaseCommController) EditItemMatchOpen(ctx echo.Context) (err error) {
	reqdata := new(ItemMatchReq)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}

	data.ItemOpenOrCloseMatch(reqdata.Id, 1)
	return c.BSucJson(ctx, reqdata)
}

// 禁用
func (c *BaseCommController) EditItemMatchClose(ctx echo.Context) (err error) {
	reqdata := new(ItemMatchReq)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.ItemOpenOrCloseMatch(reqdata.Id, 2)
	return c.BSucJson(ctx, reqdata)
}

// 机器人部分
// 获取列表
func (c *BaseCommController) GetBotSettings(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetAllBotSetting()
	count := len(resp)
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

func (c *BaseCommController) AddBotSetting(ctx echo.Context) (err error) {
	reqdata := new(mdb.BotSettingModel)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.AddBotSettingData(reqdata)
	return c.BSucJson(ctx, reqdata)
}
func (c *BaseCommController) EditBotSetting(ctx echo.Context) (err error) {
	reqdata := new(mdb.BotSettingModel)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.EditBotSettingData(reqdata)
	return c.BSucJson(ctx, reqdata)
}

type BotSetReq struct {
	Id    int `json:"id"`
	Ctype int `json:"ctype"`
}

func (c *BaseCommController) OpenorCloseBotSet(ctx echo.Context) (err error) {
	reqdata := new(BotSetReq)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	data.BotSetgingOpenOrClose(reqdata.Id, reqdata.Ctype)
	return c.BSucJson(ctx, reqdata)
}

// 轮播图
func (c *BaseCommController) SysGetImages(ctx echo.Context) (err error) {
	pagedata := new(page.BPageInfo)
	resp, _ := data.GetAllSwipers()
	count := len(resp)
	respJson := page.BPageResponse{
		PageCount: count,
		Page:      pagedata.Page,
		PageSize:  pagedata.PageSize,
		List:      resp,
	}
	return c.BSucJson(ctx, respJson)
}

// ApiMultiUpload
func (c *BaseCommController) ApiMultiUpload(ctx echo.Context) (err error) {
	//form, _ := ctx.MultipartForm()
	file, err := ctx.FormFile("files")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建文件夹
	currentTime := time.Now()
	folderName := currentTime.Format("20060102")
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100) // 生成0到99之间的随机数
	// 创建文件名，格式为"月日时分秒_随机数"
	fileRadomName := fmt.Sprintf("%02d%02d%02d%02d%02d_%02d", currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), randomNum)
	folderPath := "upload/" + folderName
	if _, errq := os.Stat(folderPath); os.IsNotExist(errq) {
		// 目标文件夹不存在，创建它
		err2 := os.MkdirAll(folderPath, os.ModePerm)
		if err2 != nil {
			fmt.Println("无法创建文件夹:", err2)
			return err2
		}
		fmt.Println("文件夹创建成功:", folderPath)
	} else {
		// 目标文件夹已存在
		fmt.Println("文件夹已存在:", folderPath)
	}

	fullPathName := folderPath + "/" + fileRadomName + file.Filename
	// 在服务器上创建一个新文件用于保存上传的图片
	dst, err := os.Create(fullPathName)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 将上传的图片复制到新文件中
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	addResp := mdb.GameSwiper{
		Imgurl: "/" + fullPathName,
		IsShow: 1,
		Sta:    1,
	}
	returnResp := data.AddSwiperData(&addResp)
	return c.BSucJson(ctx, returnResp)
}

// 删除图片
type DelImgId struct {
	ID int `json:"id"`
}

func (c *BaseCommController) DeleteImg(ctx echo.Context) (err error) {
	reqdata := new(DelImgId)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	returnResp := data.DelSwiperById(reqdata.ID)
	return c.BSucJson(ctx, returnResp)
}
