package data

import (
	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
	"github.com/assimon/luuu/util/page"
)

// INSERT INTO `grabs`.`sys_setting` (`id`, `cname`, `ckey`, `cvalue`, `mech_id`, `create_at`, `sta`, `desp`) VALUES (1, '提现配置项', 'withdrawRate', '20', 0, NULL, '1', '数值是按照提现的百分比扣除');
// 获取首页列表信息
func GetConfigLists(pageinfo *page.BPageInfo) ([]mdb.SysConfig, error) {
	var GSysconfig []mdb.SysConfig
	if pageinfo.Page <= 0 {
		pageinfo.Page = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(GSysconfig).Where("sta = ? ", 1).Limit(pageinfo.PageSize).Offset(pageinfo.Page).Find(&GSysconfig).Error
	return GSysconfig, err
}
func GetConfigListsCount() (int, error) {
	var GSysconfig []mdb.SysConfig
	err := dao.Mdb.Model(GSysconfig).Where("sta = ? ", 1).Find(&GSysconfig).Error
	return len(GSysconfig), err
}

func SysAddWebConfig(configdata *mdb.SysConfig) error {
	SysConfigs := new(mdb.SysConfig)
	err := dao.Mdb.Model(SysConfigs).Create(configdata).Error
	return err
}

// 更新
func SysUpdateConfig(configdata *mdb.SysConfig) error {
	SysConfig := new(mdb.SysConfig)
	err := dao.Mdb.Model(SysConfig).Where("id = ?", configdata.ID).Updates(map[string]interface{}{
		"cname":  configdata.Cname,
		"ckey":   configdata.Ckey,
		"cvalue": configdata.Cvalue,
		"desp":   configdata.Desp,
	}).Error
	return err
}

// 启用
func SysUpdateConfigOpen(configdata *mdb.SysConfig) error {
	SysConfig := new(mdb.SysConfig)
	err := dao.Mdb.Model(SysConfig).Where("id = ?", configdata.ID).Updates(map[string]interface{}{
		"sta": 1,
	}).Error
	return err
}

// 禁用
func SysUpdateConfigClose(configdata *mdb.SysConfig) error {
	SysConfig := new(mdb.SysConfig)
	err := dao.Mdb.Model(SysConfig).Where("id = ?", configdata.ID).Updates(map[string]interface{}{
		"sta": 2,
	}).Error
	return err
}

// 前端使用
// 根据配置名字提取配置数据
func FindConfigByConfigName(configname string) (*mdb.SysConfig, error) {
	SysConfig := new(mdb.SysConfig)
	err := dao.Mdb.Model(SysConfig).Where("ckey = ? and sta = ?", configname, 1).Find(&SysConfig).Error
	return SysConfig, err
}
