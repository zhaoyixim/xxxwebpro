package data

import (
	"strconv"

	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
	"github.com/assimon/luuu/util/page"
)

// GetOrderInfoByOrderId 通过客户订单号查询订单
// func GetDtlById() (*mdb.GameDetail, error) {
// 	GDetail := new(mdb.GameDetail)
// 	err := dao.Mdb.Model(GDetail).Limit(1).Find(GDetail, "sta = ?", 1).Error
// 	return GDetail, err
// }

// 添加发布游戏
func AddPublishItems(itemsdata *mdb.GameItemList) (*mdb.GameItemList, error) {
	GDtlModel := new(mdb.GameItemList)
	FindMember, _ := GetMemById(strconv.Itoa(itemsdata.Cuid))
	sigleBot, _ := GetSigleBotSetting()
	groupBot, _ := GetGroupBotSetting()
	itemsdata.Cavatar = FindMember.Avatar
	itemsdata.Cuname = FindMember.Cuname

	itemsdata.TgKefu = sigleBot.Cname
	itemsdata.TgGroup = groupBot.Cusername

	err := dao.Mdb.Model(GDtlModel).Create(itemsdata).Error
	if err != nil {
		return GDtlModel, err
	}
	err2 := dao.Mdb.Model(GDtlModel).Order("id desc").Limit(1).Find(&GDtlModel).Error
	ReduceCardPlus(itemsdata.ItemId, itemsdata.Cuid)
	return GDtlModel, err2
}

// 后台部分
// 获取发布列表
func GetPublishLists(pageinfo *page.BPageInfo) ([]mdb.GameItemList, error) {
	var gameItemList []mdb.GameItemList
	if pageinfo.Page <= 0 {
		pageinfo.Page = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(gameItemList).Where("sta = ? ", 1).Limit(pageinfo.PageSize).Offset(pageinfo.Page).Find(&gameItemList).Error
	return gameItemList, err
}
func GetPublishListsCount() (int, error) {
	var gameItemList []mdb.GameItemList
	err := dao.Mdb.Model(gameItemList).Where("sta = ? ", 1).Find(&gameItemList).Error
	return len(gameItemList), err
}
