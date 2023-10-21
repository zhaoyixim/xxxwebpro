package mdb

type RewordListModel struct {
	ID         int    `gorm:"column:id;primary_key" json:"id"`
	Cuid       int    `gorm:"column:cuid" json:"cuid"`
	ItemId     int    `gorm:"column:item_id" json:"item_id"`
	ItemListId int    `gorm:"column:item_list_id" json:"item_list_id"`
	RewordNum  int    `gorm:"column:reword_num" json:"reword_num"`
	Sta        int    `gorm:"column:sta" json:"sta"`
	RewordU    int    `gorm:"column:reword_U" json:"reword_U"`
	CreateAt   string `gorm:"column:create_at" json:"create_at"`
}
type ReturnTitleAndAvatar struct {
	Ctitle string `json:"ctitle"`
	Cover  string `json:"cover"`
}

func (r *RewordListModel) TableName() string {
	return "reword_list"
}
