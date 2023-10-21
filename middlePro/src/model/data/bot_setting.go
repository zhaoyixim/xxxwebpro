package data

import (
	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
)

// 获取限制数
func GetAllBotSetting() ([]mdb.BotSettingModel, error) {
	var GItem []mdb.BotSettingModel
	err := dao.Mdb.Model(GItem).Where("sta = ? ", 1).Find(&GItem).Error
	return GItem, err
}

// 添加数据
func AddBotSettingData(adddata *mdb.BotSettingModel) error {
	ItetBotModel := new(mdb.BotSettingModel)
	err := dao.Mdb.Model(ItetBotModel).Create(adddata).Error
	return err
}

// 获取群机器人的token
func GetGroupBotSetting() (*mdb.BotSettingModel, error) {
	GItem := new(mdb.BotSettingModel)
	err := dao.Mdb.Model(GItem).Limit(1).Find(&GItem, "sta = ? and ctype ", 1, 2).Error
	return GItem, err
}

// 获取系统msg机器人token
func GetSigleBotSetting() (*mdb.BotSettingModel, error) {
	GItem := new(mdb.BotSettingModel)
	err := dao.Mdb.Model(GItem).Limit(1).Find(&GItem, "sta = ? and ctype ", 1, 1).Error
	return GItem, err
}

// 编更新数据
func EditBotSettingData(editdata *mdb.BotSettingModel) error {
	ItemMatchUid := new(mdb.BotSettingModel)
	err := dao.Mdb.Model(ItemMatchUid).Where("id = ?", editdata.ID).Updates(map[string]interface{}{
		"tgtoken":   editdata.Tgtoken,
		"belongs":   editdata.Belongs,
		"cname":     editdata.Cname,
		"cusername": editdata.Cusername,
		"mech_id":   editdata.MechId,
		"ctype":     editdata.Ctype,
	}).Error
	return err
}

// 开启 禁用 1-开启 2-禁用
func BotSetgingOpenOrClose(id int, ctype int) error {
	botSettingModel := new(mdb.BotSettingModel)
	err := dao.Mdb.Model(botSettingModel).Where("id = ?", id).Updates(map[string]interface{}{
		"sta": ctype,
	}).Error
	return err
}
