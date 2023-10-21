package data

import (
	"fmt"

	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
	"gorm.io/gorm"
)

// 内部使用---- 定义一个函数来查找并返回匹配的对象
func findItemId(findArr []mdb.CuidItemModel, targetName int) *mdb.CuidItemModel {
	for _, item := range findArr {
		if item.ItemId == targetName {
			return &item
		}
	}
	return nil
}

// 前端-发布之后扣除remain_N
func ReduceCardPlus(itemid int, cuid int) error {
	GItem := new(mdb.CuidItemModel)
	err := dao.Mdb.Model(GItem).Where("item_id = ? and cuid =?", itemid, cuid).Updates(map[string]interface{}{
		"remain_N": gorm.Expr("remain_N - ?", 1),
	}).Error
	return err
}

// 检测权限
func CheckAuthByUid(reqdata *mdb.PubCheckUidItemPostData) (*mdb.CuidItemModel, error) {
	GItem := new(mdb.CuidItemModel)
	err := dao.Mdb.Model(GItem).Limit(1).Find(&GItem, "sta = ? and cuid= ? and item_id=?", 1, reqdata.Cuid, reqdata.Itemid).Error
	return GItem, err
}

// 获取限制数limitN
func GetLimitNum(itemid string, cuid int) (*mdb.CuidItemModel, error) {
	GItem := new(mdb.CuidItemModel)
	err := dao.Mdb.Model(GItem).Limit(1).Find(&GItem, "sta = ? and cuid= ? and item_id=?", 1, cuid, itemid).Error
	return GItem, err
}
func GetLimitNumByGListId(listitemid string, cuid int) (*mdb.CuidItemModel, error) {
	GItemList := new(mdb.GameItemList)
	errs := dao.Mdb.Model(GItemList).Where("id =?", listitemid).Limit(1).Find(&GItemList).Error
	if errs != nil {
		fmt.Println(errs)
	}
	GItem := new(mdb.CuidItemModel)
	err := dao.Mdb.Model(GItem).Limit(1).Find(&GItem, "sta = ? and cuid= ? and item_id=?", 1, cuid, GItemList.ItemId).Error
	return GItem, err
}

// 获取所有匹配数据
func GetAllMatchData() ([]mdb.CuidItemModel, error) {
	var GItem []mdb.CuidItemModel
	err := dao.Mdb.Model(GItem).Where("sta = ? ", 1).Find(&GItem).Error
	return GItem, err
}

// 添加数据
func AddItemMatchData(adddata *mdb.CuidItemModel) error {
	ItemMatchUid := new(mdb.CuidItemModel)
	err := dao.Mdb.Model(ItemMatchUid).Create(adddata).Error
	return err
}

// 编辑数据
func EditItemMatchData(editdata *mdb.CuidItemModel) error {
	ItemMatchUid := new(mdb.CuidItemModel)
	err := dao.Mdb.Model(ItemMatchUid).Where("id = ?", editdata.ID).Updates(map[string]interface{}{
		"cuid":       editdata.Cuid,
		"item_id":    editdata.ItemId,
		"card_N":     editdata.CardN,
		"is_limit_N": editdata.IsLimitN,
		"remain_N":   editdata.RemainN,
	}).Error
	return err
}

// 开启 禁用 1-开启 2-禁用
func ItemOpenOrCloseMatch(id int, ctype int) error {
	ItemMatchUid := new(mdb.CuidItemModel)
	err := dao.Mdb.Model(ItemMatchUid).Where("id = ?", id).Updates(map[string]interface{}{
		"sta": ctype,
	}).Error
	return err
}

// 根据UID获取所有Item
func GetItemsByuid(uid string) ([]mdb.GameItems, error) {
	var GItems []mdb.CuidItemModel
	var GameItems []mdb.GameItems
	var GItemscards []mdb.GameItems
	err := dao.Mdb.Model(GItems).Where("sta = ? AND cuid = ?", 1, uid).Find(&GItems).Error
	//获取对应的itemid
	//根据itemid 获取item信息
	if len(GItems) > 0 {
		var listID []int
		for _, item := range GItems {
			listID = append(listID, item.ID)
		}
		err := dao.Mdb.Model(GameItems).Where("sta = ? and id IN (?) ", 1, listID).Find(&GameItems).Error
		if len(GameItems) > 0 {
			for _, it := range GameItems {
				finDitem := findItemId(GItems, it.ID)
				Ditems := it
				Ditems.CardN = finDitem.CardN
				Ditems.RemainN = finDitem.RemainN
				GItemscards = append(GItemscards, Ditems)
			}
			return GItemscards, err
		}
		return GameItems, err
	}
	return GameItems, err
}
