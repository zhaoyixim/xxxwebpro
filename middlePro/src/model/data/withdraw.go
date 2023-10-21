package data

import (
	"fmt"
	"strconv"

	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
	"github.com/assimon/luuu/util/page"
)

// 添加数据
func SetInAccountResult(withdrawData *mdb.WithdrawModel) error {
	WithdrawList := new(mdb.WithdrawModel)
	configInfo, _ := FindConfigByConfigName("withdrawRate")
	if configInfo.ID > 0 {
		intValue, _ := strconv.Atoi(configInfo.Cvalue)
		withdrawData.Rate = intValue
		floatValue := float64(intValue) / 100.0
		resultVal := float64(WithdrawList.Amount) * (1 - floatValue)
		withdrawData.WithdrawU = fmt.Sprintf("%.2f", resultVal)
	}

	err := dao.Mdb.Model(WithdrawList).Create(withdrawData).Error
	return err
}

// 获取数据--付款
func GetWithdrawIn(pageinfo *page.BPageInfo) ([]mdb.WithdrawModel, error) {
	var WithdrawList []mdb.WithdrawModel
	if pageinfo.Page <= 0 {
		pageinfo.Page = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(WithdrawList).Where("ctype = ? ", 1).Limit(pageinfo.PageSize).Offset(pageinfo.Page).Find(&WithdrawList).Error
	return WithdrawList, err
}

// 获取数据
func GetWithdrawInCount() (int, error) {
	var WithdrawList []mdb.WithdrawModel
	err := dao.Mdb.Model(WithdrawList).Where("ctype = ? ", 1).Find(&WithdrawList).Error
	return len(WithdrawList), err
}

// 提现
func GetWithdrawOut(pageinfo *page.BPageInfo) ([]mdb.WithdrawModel, error) {
	var WithdrawList []mdb.WithdrawModel
	if pageinfo.Page <= 0 {
		pageinfo.Page = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(WithdrawList).Where("ctype = ? ", 2).Limit(pageinfo.PageSize).Offset(pageinfo.Page).Find(&WithdrawList).Error
	return WithdrawList, err
}
func GetWithdrawOutCount() (int, error) {
	var WithdrawList []mdb.WithdrawModel
	err := dao.Mdb.Model(WithdrawList).Where("ctype = ? ", 2).Find(&WithdrawList).Error
	return len(WithdrawList), err
}

// 提现成功
func UpdataWithDrawSta(Cit int) error {
	WithdrawList := new(mdb.WithdrawModel)
	err := dao.Mdb.Model(WithdrawList).Where("id = ?", Cit).Updates(map[string]interface{}{
		"sta": 2,
	}).Error
	return err
}

// 提现失败
func UpdateErrWithDrawSta(Cit int) error {
	WithdrawList := new(mdb.WithdrawModel)
	err := dao.Mdb.Model(WithdrawList).Where("id = ?", Cit).Updates(map[string]interface{}{
		"sta": 3,
	}).Error
	return err
}

// 提现前端删除
func DeleteWithDrawSta(wid string) error {
	WithdrawList := new(mdb.WithdrawModel)
	err := dao.Mdb.Model(WithdrawList).Where("id = ?", wid).Updates(map[string]interface{}{
		"sta": 4,
	}).Error
	return err
}
