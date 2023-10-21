package data

import (
	"fmt"
	"time"

	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
	"github.com/assimon/luuu/util/page"
)

func AddRewordRecord(record *mdb.RewordListModel) error {
	AddRecord := new(mdb.RewordListModel)
	currentTime := time.Now()
	record.CreateAt = currentTime.Format("2006-01-02 15:04:05")
	err := dao.Mdb.Model(AddRecord).Create(record).Error
	return err
}
func GetRewordRecordList(pageinfo *page.BPageInfo) ([]mdb.RewordListModel, error) {
	var rewordListModel []mdb.RewordListModel
	if pageinfo.Page <= 0 {
		pageinfo.Page = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(rewordListModel).Where("sta = ? ", 1).Limit(pageinfo.PageSize).Offset(pageinfo.Page).Find(&rewordListModel).Error
	return rewordListModel, err
}
func GetRewordRecordCount() (int, error) {
	var rewordListModel []mdb.RewordListModel
	err := dao.Mdb.Model(rewordListModel).Where("sta = ? ", 1).Find(&rewordListModel).Error
	return len(rewordListModel), err
}

func GetRewordRecordCountByUid(cuid string) (int, error) {
	var rewordListModel []mdb.RewordListModel
	err := dao.Mdb.Model(rewordListModel).Where("sta = ? and cuid=?", 1, cuid).Find(&rewordListModel).Error
	//reword_U
	totalU := 0
	for _, item := range rewordListModel {
		totalU += item.RewordU
	}
	return totalU, err
}

type ReturnTitleJson struct {
	RewordList mdb.RewordListModel
	TitleData  mdb.ReturnTitleAndAvatar
}

// 前台使用-中奖记录
func GetRewordListsByUid(cuid string, pageinfo *page.PageInfo) ([]ReturnTitleJson, error) {
	if pageinfo.CurrentPage <= 0 {
		pageinfo.CurrentPage = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	var rewordListModel []mdb.RewordListModel
	var returnJson []ReturnTitleJson
	err := dao.Mdb.Model(rewordListModel).Where("cuid = ? and sta = ? ", cuid, 1).Limit(pageinfo.PageSize).Offset(pageinfo.CurrentPage).Find(&rewordListModel).Error
	if len(rewordListModel) > 0 {
		for i, reword := range rewordListModel {
			parsedTime, err := time.Parse(time.RFC3339, reword.CreateAt)
			if err != nil {
				fmt.Println("解析错误:", err)
				continue
			}
			rewordListModel[i].CreateAt = parsedTime.Format("2006-01-02 15:04:05")
			findTitleAndAvatar := GetTitleById(reword.ItemId)
			returnData := ReturnTitleJson{
				RewordList: rewordListModel[i],
				TitleData:  *findTitleAndAvatar,
			}
			returnJson = append(returnJson, returnData)
		}
	}
	return returnJson, err
}

// 前台使用-提现记录
func GetWithDrawListsByUid(cuid string, pageinfo *page.PageInfo) ([]mdb.WithdrawModel, error) {
	var withDrawListModel []mdb.WithdrawModel
	if pageinfo.CurrentPage <= 0 {
		pageinfo.CurrentPage = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(withDrawListModel).Where("cuid = ? and sta < ? ", cuid, 4).Limit(pageinfo.PageSize).Offset(pageinfo.CurrentPage).Find(&withDrawListModel).Error
	return withDrawListModel, err
}

// 中奖前端删除
func DeleteRewordSta(wid string) error {
	RewordModel := new(mdb.RewordListModel)
	err := dao.Mdb.Model(RewordModel).Where("id = ?", wid).Updates(map[string]interface{}{
		"sta": 2,
	}).Error
	return err
}
