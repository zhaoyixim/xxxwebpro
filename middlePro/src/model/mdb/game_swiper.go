package mdb

type GameSwiper struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	Imgurl string `gorm:"column:imgurl" json:"imgurl"`
	IsShow int    `gorm:"column:is_show" json:"is_show"`
	Sta    int    `gorm:"column:sta" json:"sta"`
}

func (w *GameSwiper) TableName() string {
	return "game_swiper"
}
