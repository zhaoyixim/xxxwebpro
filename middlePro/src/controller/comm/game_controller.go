package comm

import (
	"github.com/labstack/echo/v4"
	"github.com/pengk/summer/model/data"
	"github.com/pengk/summer/model/mdb"
	"github.com/pengk/summer/util/page"
)

// 首页信息
func (c *BaseCommController) IndexGLists(ctx echo.Context) (err error) {
	resp, err := data.GetIndexLists()
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
	// fmt.Println(resp)

	// tmpl, err := template.ParseFiles(fmt.Sprintf(".%s/%s", config.StaticPath, "index.html"))
	// if err != nil {
	// 	return ctx.String(http.StatusOK, err.Error())
	// }
	// return tmpl.Execute(ctx.Response(), resp)
}

// 获取我能够发布的游戏列表
func (c *BaseCommController) IndexGMyList(ctx echo.Context) (err error) {
	Cuid := ctx.Param("uid")
	//uid, _ = strconv.Atoi(Cuid)
	resp, err := data.GetItemsByuid(Cuid)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 检查添加发布时候，是否有权限
func (c *BaseCommController) CheckIsPublish(ctx echo.Context) (err error) {
	reqdata := new(mdb.PubCheckUidItemPostData)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	//PubPostData

	resp, err := data.CheckAuthByUid(reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 轮播图
func (c *BaseCommController) IndexGSwiper(ctx echo.Context) (err error) {
	resp, err := data.GetSwiperLists()
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 通过ID获取游戏列表-进行中
func (c *BaseCommController) GetUnfinishedGList(ctx echo.Context) (err error) {
	Listid := ctx.Param("tid")
	reqdata := new(page.PageInfo)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	resp, err := data.GetGUnfinishedListById(Listid, reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 通过ID获取游戏列表-已完成
func (c *BaseCommController) GetfinishedGList(ctx echo.Context) (err error) {
	Listid := ctx.Param("tid")
	reqdata := new(page.PageInfo)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	resp, err := data.GetGFinishedListById(Listid, reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

/*游戏发布记录*/
func (c *BaseCommController) GetPublishfinished(ctx echo.Context) (err error) {
	Listid := ctx.Param("uid")
	reqdata := new(page.PageInfo)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}
	resp, err := data.GetPublishedfinishedList(Listid, reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}
func (c *BaseCommController) GetPublishUnfinished(ctx echo.Context) (err error) {
	Listid := ctx.Param("uid")
	reqdata := new(page.PageInfo)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}

	resp, err := data.GetPublishedUnfinishedList(Listid, reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 中间菜单列表----进行中的所有列表
func (c *BaseCommController) GetAllDodingG(ctx echo.Context) (err error) {
	reqdata := new(page.PageInfo)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}

	resp, err := data.GetAllunFinishingData(reqdata)
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 游戏详情
type GetGDtlReqM struct {
	Tid  int `json:"tid"`
	Cuid int `json:"cuid"`
}
type GetGDtlResponse struct {
	GTitem          *mdb.GameItemList `json:"itemdata"`
	RewordOrderItem []mdb.Orders      `json:"rewordorder"`
	LimitType       int               `json:"is_limit_N"`
}

func (c *BaseCommController) GetGDtl(ctx echo.Context) (err error) {
	reqdata := new(GetGDtlReqM)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}

	//p, _ := data.GetGDtlById(strconv.Itoa(reqdata.Tid))
	p, s, limitN, err := data.GetIsRewordAndData(reqdata.Tid, reqdata.Cuid)

	resp := GetGDtlResponse{
		GTitem:          p,
		RewordOrderItem: s,
		LimitType:       limitN,
	}
	if err != nil {
		return c.FailJson(ctx, err)
	}
	return c.SucJson(ctx, resp)
}

// 发布游戏
func (c *BaseCommController) PubGame(ctx echo.Context) (err error) {
	reqdata := new(mdb.GameItemList)
	if err = ctx.Bind(reqdata); err != nil {
		return c.FailJson(ctx, err)
	}

	findItem, errs := data.AddPublishItems(reqdata)
	if errs != nil {
		return c.FailJson(ctx, errs)
	}

	return c.SucJson(ctx, findItem)
	// returnjson := "添加成功"
	// return c.SuccJFailson(ctx, returnjson)
}

// 获取配置
type CgByName struct {
	Ckey string `json:"ckey"`
}

func (c *BaseCommController) GetConfigByName(ctx echo.Context) (err error) {
	requestData := new(CgByName)
	if err = ctx.Bind(requestData); err != nil {
		return c.FailJson(ctx, err)
	}
	sysinfo, err2 := data.FindConfigByConfigName(requestData.Ckey)
	if err != nil {
		return c.FailJson(ctx, err2)
	}
	return c.SucJson(ctx, sysinfo)
}
