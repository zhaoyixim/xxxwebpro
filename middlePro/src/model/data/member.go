package data

import (
	"fmt"
	"time"

	"github.com/pengk/summer/model/dao"
	"github.com/pengk/summer/model/mdb"
	"github.com/pengk/summer/util/page"
	"gorm.io/gorm"
)

// 通过ID查member
func GetMemById(Mid string) (*mdb.MemberModel, error) {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Limit(1).Find(GMember, "sta = ? and id= ?", 1, Mid).Error
	return GMember, err
}

// 添加钱包地址
func EditMemberWalletAddress(cuid int, addresswallet string) (*mdb.MemberModel, error) {
	GMember := new(mdb.MemberModel)
	fmt.Println(cuid)
	err2 := dao.Mdb.Model(GMember).Where("id = ?", cuid).Updates(map[string]interface{}{
		"wallet_address": addresswallet,
	}).Error
	if err2 != nil {
		return GMember, err2
	}

	err := dao.Mdb.Model(GMember).Where("id = ? ", cuid).Find(&GMember).Error
	return GMember, err
}

// 查出所有member
func GetAllMember(Gid string) ([]mdb.MemberModel, error) {
	var GMembers []mdb.MemberModel
	err := dao.Mdb.Model(GMembers).Where("sta = ? ", 2).Find(&GMembers).Error
	return GMembers, err
}

// 通过USERNAME查找
func FindByUseranme(cuname string) (*mdb.MemberModel, error) {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Limit(1).Find(GMember, "cuname = ?", cuname).Error
	if err != nil {
		return GMember, err
	}
	return GMember, nil
}

// 注册
func RegAndAddmember(memberdata *mdb.MemberModel) error {
	Addmembers := new(mdb.MemberModel)
	err := dao.Mdb.Model(Addmembers).Create(memberdata).Error
	return err
}

// 登录
func LoginAndUpdata(mid int) error {
	GMember := new(mdb.MemberModel)
	currentTime := time.Now()
	mysqlTimeFormat := currentTime.Format("2006-01-02 15:04:05")
	err2 := dao.Mdb.Model(GMember).Where("id = ?", mid).Updates(map[string]interface{}{
		"update_at": mysqlTimeFormat,
		"sta":       1,
	}).UpdateColumn("login_times", gorm.Expr("login_times + ?", 1)).Error
	return err2
}

// 设置登录过期--减少系统轮询资源
func SetExpireLogin(mid string) error {
	GMember := new(mdb.MemberModel)
	currentTime := time.Now()
	mysqlTimeFormat := currentTime.Format("2006-01-02 15:04:05")
	err2 := dao.Mdb.Model(GMember).Where("id = ?", mid).Updates(map[string]interface{}{
		"sta":       3,
		"update_at": mysqlTimeFormat,
	}).Error
	return err2
}

// 更新我的钱包信息WalletAddress
func UpdateMyWalletAddress(wal_address string, mid int) error {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Where("id = ?", mid).Updates(map[string]interface{}{
		"wallet_address": wal_address,
	}).Error
	return err
}

/*查询付款钱包*/
//轮询钱包地址
//ListenTrc20Job 使用
func GetAvailableMemberWallet() ([]mdb.MemberModel, error) {
	// 获取当前时间
	currentTime := time.Now()
	// 计算 15 分钟前的时间
	fifteenMinutesAgo := currentTime.Add(-15 * time.Minute)
	var GMembers []mdb.MemberModel
	//15分钟内
	err := dao.Mdb.Model(GMembers).Where("update_at < ? and sta = ?", fifteenMinutesAgo, 1).Find(&GMembers).Error
	if err != nil {
		fmt.Println("查询数据时出错:", err)
		return GMembers, err
	}
	return GMembers, err
}

//上面方法结果整理数组

func RelizeArrMember() (map[string]mdb.MemberModel, error) {
	findmebersArr, err := GetAvailableMemberWallet()
	resps := make(map[string]mdb.MemberModel)
	for _, item := range findmebersArr {
		if item.WalletAddress != "" {
			resps[item.WalletAddress] = item
		}
	}
	return resps, err
}

// 添加到额外的U
func UpdateToExtraU(extraU int, mid int, itemId string) error {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Where("id = ?", mid).UpdateColumn("extra_U", gorm.Expr("extra_U + ?", extraU)).Error
	GDtlModel := new(mdb.GameItemList)
	currentTime := time.Now()
	dao.Mdb.Model(GDtlModel).Where("id = ?", itemId).Updates(map[string]interface{}{
		"paying_num": gorm.Expr("paying_num + ?", 1),
		"updated_at": currentTime.Format("2006-01-02 15:04:05"),
	})

	return err
}
func ReduceToExtraU(reduceExtraU int, mid int) error {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Where("id = ?", mid).UpdateColumn("extra_U", gorm.Expr("extra_U - ?", reduceExtraU)).Error
	return err
}

// 添加到total_U
func UpdateRewordToTotalU(addRewordU int, mid int) error {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Where("id = ?", mid).UpdateColumn("total_U", gorm.Expr("total_U + ?", addRewordU)).Error

	return err
}

//提现

func ReduceRewordTotalU(reduceRewordU int, mid int) error {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Where("id = ?", mid).UpdateColumn("total_U", gorm.Expr("total_U - ?", reduceRewordU)).Error
	return err
}

//后台部分--获取所有用户信息

func BGetAllUserLists(pageinfo *page.BPageInfo) ([]mdb.MemberModel, error) {
	var memberList []mdb.MemberModel
	if pageinfo.Page <= 0 {
		pageinfo.Page = 0
	}
	if pageinfo.PageSize <= 0 {
		pageinfo.PageSize = 10
	}
	err := dao.Mdb.Model(memberList).Where("sta = ? ", 1).Limit(pageinfo.PageSize).Offset(pageinfo.Page).Find(&memberList).Error
	return memberList, err
}
func BGetAllUserCount() (int, error) {
	var memberListCount []mdb.MemberModel
	err := dao.Mdb.Model(memberListCount).Where("sta = ? ", 1).Find(&memberListCount).Error
	return len(memberListCount), err
}

//更改用户商户身份

func ChangeMemMechId(cuid int) error {
	GMember := new(mdb.MemberModel)
	err := dao.Mdb.Model(GMember).Where("id = ?", cuid).Updates(map[string]interface{}{
		"is_mech": 1,
	}).Error
	return err
}

/*后台首页*/
//总注册数
type GetbackIndexVisits struct {
	AllVisits float64 `json:"allvisits"`
	DayVisits float64 `json:"dayvisits"`
	DayActive float64 `json:"dayactive"`
}

func BackGetRegVisits() (*GetbackIndexVisits, error) {
	var GMembers []mdb.MemberModel
	var GMember []mdb.MemberModel
	var GMemberTodayUpdate []mdb.MemberModel
	errs := dao.Mdb.Model(GMembers).Where("sta = ?", 1).Find(&GMembers).Error
	if errs != nil {
		return nil, errs
	}
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if err := dao.Mdb.Model(GMember).Where("sta = ? and create_at >= ?", 1, today.Format("2006-01-02 15:04:05")).Find(&GMember).Error; err != nil {
		return nil, err
	}

	if err2 := dao.Mdb.Model(GMemberTodayUpdate).Where("sta = ? and update_at >= ?", 1, today.Format("2006-01-02 15:04:05")).Find(&GMemberTodayUpdate).Error; err2 != nil {
		return nil, err2
	}

	resp := &GetbackIndexVisits{
		AllVisits: float64(len(GMembers)),
		DayVisits: float64(len(GMember)),
		DayActive: float64(len(GMemberTodayUpdate)),
	}
	return resp, nil
}

// 前端更改用户信息
func EditMemInfoModel(memberinfo *mdb.MemberModel) (*mdb.MemberModel, error) {
	GMember := new(mdb.MemberModel)
	errs := dao.Mdb.Model(GMember).Where("id = ?", memberinfo.ID).Updates(map[string]interface{}{
		"nickname": memberinfo.Nickname,
		"avatar":   memberinfo.Avatar,
	}).Error
	if errs != nil {
		return nil, errs
	}
	return memberinfo, nil
}
