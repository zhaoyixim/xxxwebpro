package data

import (
	"strconv"

	"github.com/golang-module/carbon/v2"
	"github.com/pengk/summer/model/dao"
	"github.com/pengk/summer/model/mdb"
)

// 生成MSGCODE信息
func GenerateMsgCode(codedata *mdb.MsgCode) error {
	msgcode := new(mdb.MsgCode)
	err := dao.Mdb.Model(msgcode).Create(codedata).Error
	ClearMsgExpire()
	return err
}

// 查询msgcode
func FindMsgCodeByCode(msgcode string) (*mdb.MsgCode, error) {
	MsgCodeModel := new(mdb.MsgCode)
	err := dao.Mdb.Model(MsgCodeModel).Limit(1).Find(MsgCodeModel, "msgcode = ?", msgcode).Error
	return MsgCodeModel, err
}

// 更新msgcode---已使用
func UpdataMsgCodeByCode(msgcode string) (*mdb.MsgCode, error) {
	MsgCodeUpdata := new(mdb.MsgCode)
	//更新msg表
	err2 := dao.Mdb.Model(MsgCodeUpdata).Where("msgcode = ?", msgcode).Updates(map[string]interface{}{
		"sta": 2,
	}).Error
	return MsgCodeUpdata, err2
}

func ClearMsgExpire() ([]mdb.MsgCode, error) {
	var MsgLists []mdb.MsgCode
	err := dao.Mdb.Model(MsgLists).Where("sta = ?", 0).Find(&MsgLists).Error
	currentTime := carbon.Now().ToDateTimeString()
	//5分钟有效
	var fiveStableSeconds int64 = 10 * 60
	for _, listItem := range MsgLists {
		msgcode, _ := strconv.Atoi(listItem.Msgcode)
		if msgcode > 0 && listItem.Sta == 0 {
			var timesting = carbon.Parse(listItem.CreateAt.String()).DiffInSeconds(carbon.Parse(currentTime))
			if fiveStableSeconds <= timesting {
				//过期了
				//fmt.Println(listItem.ID)
				MsgErr := dao.Mdb.Model(MsgLists).Where("id = ?", listItem.ID).Updates(map[string]interface{}{
					"sta": 3,
				}).Error
				if MsgErr != nil {
					return MsgLists, MsgErr
				}
			}
		}

	}
	return MsgLists, err
}
