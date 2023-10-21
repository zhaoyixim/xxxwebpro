package data

import (
	"fmt"

	"github.com/assimon/luuu/model/dao"
	"github.com/assimon/luuu/model/mdb"
)

func GetAllSwipers() ([]mdb.GameSwiper, error) {
	var GSwiper []mdb.GameSwiper
	err := dao.Mdb.Model(GSwiper).Where("sta = ? ", 1).Find(&GSwiper).Error
	return GSwiper, err
}

// 添加数据
func AddSwiperData(adddata *mdb.GameSwiper) *mdb.GameSwiper {
	SwiperModel := new(mdb.GameSwiper)
	SwiperReturnModel := new(mdb.GameSwiper)
	result := dao.Mdb.Model(SwiperModel).Create(adddata)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	result.Scan(&SwiperReturnModel)
	return SwiperReturnModel
}

//删除数据

func DelSwiperById(ID int) *mdb.GameSwiper {
	SwiperModel := new(mdb.GameSwiper)
	SwiperReturnModel := new(mdb.GameSwiper)
	result := dao.Mdb.Model(SwiperModel).Where("id = ?", ID).Updates(map[string]interface{}{
		"sta":     2,
		"is_show": 2,
	})
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	result.Scan(&SwiperReturnModel)
	return SwiperReturnModel
}
