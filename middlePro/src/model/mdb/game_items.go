package mdb

type GameItems struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Cover   string `gorm:"column:cover" json:"cover"`
	Icon    string `gorm:"column:icon" json:"icon"`
	Url     string `gorm:"column:url" json:"url"`
	Path    string `gorm:"column:path" json:"path"`
	Sta     int    `gorm:"column:sta" json:"sta"`
	IsShow  int    `gorm:"column:is_show" json:"is_show"`
	Desp    string `gorm:"column:desp;omitempty" json:"desp"`
	CardN   int    `json:"card_n" gorm:"column:card_n;omitempty"`
	RemainN int    `json:"remain_N" gorm:"column:card_n;omitempty"`
}

// TableName sets the insert table name for this struct type
func (w *GameItems) TableName() string {
	return "game_items"
}
