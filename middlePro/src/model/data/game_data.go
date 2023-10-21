package data

import (
	"fmt"
	"strconv"
	"time"

	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
	"github.com/assimon/luuu/util/page"
	"gorm.io/gorm"
)

// 获取首页列表信息
func GetIndexLists() ([]mdb.GameItems, error) {
	var GItemList []mdb.GameItems
	err := dao.Mdb.Model(GItemList).Where("sta = ? AND is_show = ?", 1, 1).Find(&GItemList).Error
	return GItemList, err
}

// h获取我能够发布的游戏列表 --需要改
// func GetIndexListByUid(uid string) (*mdb.GameItems, error) {
// 	var GItemLists = new(mdb.GameItemList)
// 	err := dao.Mdb.Model(GItemLists).Limit(1).Find(&GItemLists, "cuid = ? ", uid).Error
// 	return GItemList, err
// }

// 获取轮播图
func GetSwiperLists() ([]mdb.GameSwiper, error) {
	var GSwiperList []mdb.GameSwiper
	err := dao.Mdb.Model(GSwiperList).Where("sta = ? AND is_show = ?", 1, 1).Find(&GSwiperList).Error
	return GSwiperList, err
}

// 通过ID获取游戏列表-进行中
func GetGUnfinishedListById(Gid string, pageinfo *page.PageInfo) ([]mdb.GameItemList, error) {
	var GItemLists []mdb.GameItemList
	if pageinfo.CurrentPage <= 1 {
		pageinfo.CurrentPage = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(GItemLists).Where("item_id = ? and sta = ? ", Gid, 1).Limit(pageinfo.PageSize).Offset(pageinfo.CurrentPage).Find(&GItemLists).Error
	return GItemLists, err
}

// 通过ID获取游戏列表-已完成
func GetGFinishedListById(Gid string, pageinfo *page.PageInfo) ([]mdb.GameItemList, error) {
	var GItemLists []mdb.GameItemList
	if pageinfo.CurrentPage <= 1 {
		pageinfo.CurrentPage = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(GItemLists).Where("item_id = ? and sta = ? ", Gid, 2).Limit(pageinfo.PageSize).Offset(pageinfo.CurrentPage).Find(&GItemLists).Error
	return GItemLists, err
}

/*游戏发布记录*/
func GetPublishedfinishedList(Uid string, pageinfo *page.PageInfo) ([]mdb.GameItemList, error) {
	var GItemLists []mdb.GameItemList
	if pageinfo.CurrentPage <= 1 {
		pageinfo.CurrentPage = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(GItemLists).Where("cuid = ? and sta = ? ", Uid, 2).Limit(pageinfo.PageSize).Offset(pageinfo.CurrentPage).Find(&GItemLists).Error
	return GItemLists, err
}
func GetPublishedUnfinishedList(Uid string, pageinfo *page.PageInfo) ([]mdb.GameItemList, error) {
	var GItemLists []mdb.GameItemList
	if pageinfo.CurrentPage <= 1 {
		pageinfo.CurrentPage = 1
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(GItemLists).Where("cuid = ? and sta = ? ", Uid, 1).Limit(pageinfo.PageSize).Offset(pageinfo.CurrentPage).Find(&GItemLists).Error
	return GItemLists, err
}

// 中间菜单列表----进行中的所有列表
func GetAllunFinishingData(pageinfo *page.PageInfo) ([]mdb.GameItemList, error) {
	var GItemLists []mdb.GameItemList
	if pageinfo.CurrentPage <= 1 {
		pageinfo.CurrentPage = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(GItemLists).Where(" sta = ? ", 1).Limit(pageinfo.PageSize).Offset(pageinfo.CurrentPage).Find(&GItemLists).Error
	return GItemLists, err
}

// 游戏详情
func GetGDtlById(Gid string) (*mdb.GameItemList, error) {
	var GItemLists = new(mdb.GameItemList)
	err := dao.Mdb.Model(GItemLists).Limit(1).Find(&GItemLists, "id = ? ", Gid).Error
	return GItemLists, err
}

// 检查是否中奖，并且中奖多少
func GetIsRewordAndData(itemid int, cuid int) (*mdb.GameItemList, []mdb.Orders, int, error) {
	GItemList := new(mdb.GameItemList)
	errs := dao.Mdb.Model(GItemList).Where("id =?", itemid).Limit(1).Find(&GItemList).Error
	if errs != nil {
		fmt.Println(errs)
	}
	var GRewordOrders []mdb.Orders
	err := dao.Mdb.Model(GRewordOrders).Where("cuid = ? and item_id = ?", cuid, itemid).Order("id ASC").Find(&GRewordOrders).Error

	limitN, _ := GetLimitNum(strconv.Itoa(GItemList.ItemId), cuid)

	return GItemList, GRewordOrders, limitN.IsLimitN, err
}

// 付款后游戏参数改变

func DealGameReword(order *mdb.Orders, GitemDtl *mdb.GameItemList, orderLen int) (*mdb.GameItemList, error) {
	GItem := new(mdb.GameItemList)
	GItemReturn := new(mdb.GameItemList)

	IsReword := 0
	orderSortNum := orderLen
	sta := 1
	reword_all := GitemDtl.RewordAll
	if GitemDtl.Reword1 > 0 {
		if GitemDtl.RewordNum1 == orderSortNum {
			IsReword = 1
			reword_all = reword_all + "第1次中奖号:" + strconv.Itoa(GitemDtl.RewordNum1) + ",(" + strconv.Itoa(GitemDtl.Reword1) + ")<br>"
			UpdateRewordToTotalU(GitemDtl.PayUnit, GitemDtl.Cuid)
			UpdateOrderIsRewordFlag(order.ID)
			//添加中奖记录
			InsertReword := &mdb.RewordListModel{
				Cuid:       order.Cuid,
				RewordNum:  GitemDtl.RewordNum1,
				RewordU:    GitemDtl.Reword1,
				ItemId:     GitemDtl.ItemId,
				ItemListId: GitemDtl.ID,
			}
			AddRewordRecord(InsertReword)
		}
		if GitemDtl.RewordNum2 == orderSortNum {
			IsReword = 1
			reword_all = reword_all + "第2次中奖号:" + strconv.Itoa(GitemDtl.RewordNum2) + ",(" + strconv.Itoa(GitemDtl.Reword2) + ")<br>"
			UpdateRewordToTotalU(GitemDtl.PayUnit, GitemDtl.Cuid)
			UpdateOrderIsRewordFlag(order.ID)

			InsertReword := &mdb.RewordListModel{
				Cuid:       order.Cuid,
				RewordNum:  GitemDtl.RewordNum2,
				RewordU:    GitemDtl.Reword2,
				ItemId:     GitemDtl.ItemId,
				ItemListId: GitemDtl.ID,
			}
			AddRewordRecord(InsertReword)
		}
		if GitemDtl.RewordNum3 == orderSortNum {
			IsReword = 1
			reword_all = reword_all + "第3次中奖号:" + strconv.Itoa(GitemDtl.RewordNum3) + ",(" + strconv.Itoa(GitemDtl.Reword3) + ")<br>"
			UpdateRewordToTotalU(GitemDtl.PayUnit, GitemDtl.Cuid)
			UpdateOrderIsRewordFlag(order.ID)

			InsertReword := &mdb.RewordListModel{
				Cuid:       order.Cuid,
				RewordNum:  GitemDtl.RewordNum3,
				RewordU:    GitemDtl.Reword3,
				ItemId:     GitemDtl.ItemId,
				ItemListId: GitemDtl.ID,
			}
			AddRewordRecord(InsertReword)
		}
		if GitemDtl.RewordNum4 == orderSortNum {
			IsReword = 1
			reword_all = reword_all + "第4次中奖号:" + strconv.Itoa(GitemDtl.RewordNum4) + ",(" + strconv.Itoa(GitemDtl.Reword4) + ")<br>"
			UpdateRewordToTotalU(GitemDtl.PayUnit, GitemDtl.Cuid)
			UpdateOrderIsRewordFlag(order.ID)

			InsertReword := &mdb.RewordListModel{
				Cuid:       order.Cuid,
				RewordNum:  GitemDtl.RewordNum4,
				RewordU:    GitemDtl.Reword4,
				ItemId:     GitemDtl.ItemId,
				ItemListId: GitemDtl.ID,
			}
			AddRewordRecord(InsertReword)
		}
		if GitemDtl.RewordNum5 == orderSortNum {
			IsReword = 1
			reword_all = reword_all + "第5次中奖号:" + strconv.Itoa(GitemDtl.RewordNum5) + ",(" + strconv.Itoa(GitemDtl.Reword5) + ")<br>"
			UpdateRewordToTotalU(GitemDtl.PayUnit, GitemDtl.Cuid)
			UpdateOrderIsRewordFlag(order.ID)

			InsertReword := &mdb.RewordListModel{
				Cuid:       order.Cuid,
				RewordNum:  GitemDtl.RewordNum5,
				RewordU:    GitemDtl.Reword5,
				ItemId:     GitemDtl.ItemId,
				ItemListId: GitemDtl.ID,
			}
			AddRewordRecord(InsertReword)
		}

	}
	findLimitN, _ := GetLimitNum(strconv.Itoa(GitemDtl.ItemId), order.Cuid)
	totallPayedU := orderSortNum * GitemDtl.PayUnit
	if findLimitN.IsLimitN == 0 {
		//第一种
		if totallPayedU >= GitemDtl.TotalU {
			sta = 2
		}
	}
	//第二种不用管-手动停止
	if findLimitN.IsLimitN == 2 {
		//第三种
		if orderSortNum >= GitemDtl.PayCount {
			sta = 2
		}
	}
	//"paying_num": orderSortNum,
	currentTime := time.Now()
	result := dao.Mdb.Model(GItem).Where("id = ?", order.ItemId).Updates(map[string]interface{}{
		"is_reword":  IsReword,
		"paying_num": gorm.Expr("paying_num + ?", 1),
		"charged_U":  gorm.Expr("charged_U + ?", GitemDtl.PayUnit),
		"sta":        sta,
		"reword_all": reword_all,
		"updated_at": currentTime.Format("2006-01-02 15:04:05"),
	})
	result.Scan(&GItemReturn)
	return GItemReturn, result.Error
}

func GetTitleById(itemid int) *mdb.ReturnTitleAndAvatar {
	GDtlModel := new(mdb.GameItemList)
	GItem := new(mdb.GameItems)
	dao.Mdb.Model(GDtlModel).Where("id = ? ", itemid).Limit(1).Find(&GDtlModel)
	dao.Mdb.Model(GItem).Where("id = ? ", GDtlModel.ItemId).Limit(1).Find(&GItem)
	returnJson := &mdb.ReturnTitleAndAvatar{
		Ctitle: GDtlModel.Ctitle,
		Cover:  GItem.Cover,
	}
	return returnJson
}

// 类别列表
// 后台
func GetItemClassLists(pageinfo *page.BPageInfo) ([]mdb.GameItems, error) {
	var GItemList []mdb.GameItems
	if pageinfo.Page <= 0 {
		pageinfo.Page = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(GItemList).Where("sta = ? ", 1).Limit(pageinfo.PageSize).Offset(pageinfo.Page).Find(&GItemList).Error
	return GItemList, err
}

func GetItemClassListsCount() (int, error) {
	var GItemList []mdb.GameItems
	err := dao.Mdb.Model(GItemList).Where("sta = ? ", 1).Find(&GItemList).Error
	return len(GItemList), err
}

// 类别操作展示
func OperaItemShow(wid int) error {
	RewordModel := new(mdb.RewordListModel)
	err := dao.Mdb.Model(RewordModel).Where("id = ?", wid).Updates(map[string]interface{}{
		"is_show": 1,
	}).Error
	return err
}

// 添加
func AddSysItemsData(adddata *mdb.GameItems) error {
	GameItem := new(mdb.GameItems)
	err := dao.Mdb.Model(GameItem).Create(adddata).Error
	return err
}

// 编辑
func EditSysItemsData(editdata *mdb.GameItems) error {
	GameItem := new(mdb.GameItems)
	err := dao.Mdb.Model(GameItem).Where("id = ?", editdata.ID).Updates(map[string]interface{}{
		"cover": editdata.Cover,
		"title": editdata.Title,
		"icon":  editdata.Icon,
		"url":   editdata.Url,
		"path":  editdata.Path,
		"desp":  editdata.Desp,
	}).Error

	return err
}

// 删除
func DeleteSysItemsData(delId int) error {
	GameItem := new(mdb.GameItems)
	err := dao.Mdb.Model(GameItem).Where("id = ?", delId).Updates(map[string]interface{}{
		"sta": 2,
	}).Error
	return err
}
